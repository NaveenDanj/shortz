package app

import (
	"authclient"
	"config"
	apiContract "contract/api"
	"fmt"
	"net/http"
	"users/internal/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) (*mux.Router, error) {

	verifier, err := authclient.NewVerifier(authclient.Config{
		JWKSURL:  cfg.UserConfig.AuthJWKSURL,
		Issuer:   cfg.UserConfig.AuthIssuer,
		Audience: cfg.UserConfig.AuthIssuer,
	})

	if err != nil {
		return nil, fmt.Errorf("auth verifier: %w", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		apiContract.ResponseWriter(w, http.StatusOK, "Users:OK")
	}).Methods(http.MethodGet)

	userController := users.NewController(users.NewService(users.NewRepository(db)))
	users.RegisterRoutes(router, userController, verifier.Middleware)

	return router, nil
}
