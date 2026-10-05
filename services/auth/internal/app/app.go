package app

import (
	"net/http"

	authfeature "auth/internal/auth"
	apiContract "contract/api"

	godevauth "github.com/go-dev-auth/go-dev-auth"

	"github.com/gorilla/mux"
)

func NewRouter(auth *godevauth.Auth) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		apiContract.ResponseWriter(w, http.StatusOK, "Auth:OK")
	}).Methods(http.MethodGet)

	authfeature.RegisterRoutes(r, auth)

	r.PathPrefix(auth.Config().BasePath).Handler(auth.Handler())

	return r
}
