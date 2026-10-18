package users

import (
	"authclient"
	apiContract "contract/api"
	"errors"
	"log"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetMe(w http.ResponseWriter, r *http.Request) {

	claims, ok := authclient.ClaimsFromContext(r.Context())
	if !ok || claims.Subject == "" {
		apiContract.ResponseWriter(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := c.service.GetMe(r.Context(), claims.Subject)
	if errors.Is(err, ErrNotFound) {
		apiContract.ResponseWriter(w, http.StatusNotFound, "Profile not found")
		return
	}

	if err != nil {
		log.Printf("users: get me %s: %v", claims.Subject, err)
		apiContract.ResponseWriter(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	apiContract.ResponseWithDataWriter(w, http.StatusOK, user)
}
