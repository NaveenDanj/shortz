package app

import (
	"encoding/json"
	"net/http"

	authfeature "auth/internal/auth"
	godevauth "github.com/go-dev-auth/go-dev-auth"

	"github.com/gorilla/mux"
)

func NewRouter(auth *godevauth.Auth) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods(http.MethodGet)

	authfeature.RegisterRoutes(r, auth)

	r.PathPrefix(auth.Config().BasePath).Handler(auth.Handler())

	return r
}
