package app

import (
	"config"
	apiContract "contract/api"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		apiContract.ResponseWriter(w, http.StatusOK, "Users:OK")
	}).Methods(http.MethodGet)

	return router
}
