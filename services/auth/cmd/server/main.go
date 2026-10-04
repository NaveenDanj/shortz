package main

import (
	"config"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	authcontract "contract/api/auth"

	godevauth "github.com/go-dev-auth/go-dev-auth"
	"github.com/gorilla/mux"

	"github.com/go-dev-auth/go-dev-auth/storage/sqlstore"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(config.AuthConfig.DatabaseDriver, config.AuthConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	store := sqlstore.New(
		db,
		sqlstore.Postgres,
		nil,
	)

	auth, err := godevauth.New(godevauth.Config{
		BaseURL:  config.AuthConfig.BaseURL,
		Secret:   config.AuthConfig.Secret,
		Database: store,

		EmailAndPassword: godevauth.EmailPasswordConfig{
			Enabled:           true,
			MinPasswordLength: authcontract.Auth_Constants.MIN_PASSWORD_LENGTH,
			MaxPasswordLength: authcontract.Auth_Constants.MAX_PASSWORD_LENGTH,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	store.SetSchema(auth.Schema())

	defer auth.StartCleanup(context.Background(), time.Hour)()

	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
