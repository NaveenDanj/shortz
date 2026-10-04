package auth

import (
	"net/http"

	godevauth "github.com/go-dev-auth/go-dev-auth"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, a *godevauth.Auth) {
	protected := r.NewRoute().Subrouter()
	protected.Use(RequireAuth(a))

	protected.HandleFunc("/me", Me).Methods(http.MethodGet)
}
