package httpx

import (
	apiContract "contract/api"
	apiUsers "contract/api/users"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const maxBodyBytes = 1 << 20 // 1 MiB

type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

var validate = newValidator()

func newValidator() *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return v
}

func DecodeAndValidate[T any](w http.ResponseWriter, r *http.Request) (*T, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var v T
	if err := dec.Decode(&v); err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			apiContract.ResponseWriter(w, http.StatusBadRequest, "Request body too large")
		} else {
			apiContract.ResponseWriter(w, http.StatusBadRequest, "Invalid request body")
		}
		return nil, false
	}

	if err := validate.Struct(&v); err != nil {
		var verrs validator.ValidationErrors
		if errors.As(err, &verrs) {
			fields := make([]FieldError, 0, len(verrs))
			for _, fe := range verrs {
				fields = append(fields, FieldError{Field: fe.Namespace(), Rule: fe.Tag()})
			}
			apiContract.ResponseWithDataWriter(w, http.StatusUnprocessableEntity, fields)
		} else {
			apiContract.ResponseWriter(w, http.StatusBadRequest, "Invalid request body")
		}
		return nil, false
	}

	return &v, true
}

func WriteError(w http.ResponseWriter, op, userID string, err error) {
	if errors.Is(err, apiUsers.ErrNotFound) {
		apiContract.ResponseWriter(w, http.StatusNotFound, "Profile not found")
		return
	}

	apiContract.ResponseWriter(w, http.StatusInternalServerError, "Internal server error")
}
