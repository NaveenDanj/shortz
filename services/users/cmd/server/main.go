package main

import (
	"config"
	"log"
	"net/http"
	"users/internal/app"
	"users/pkg/db"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	if err := app.RunMigrations(cfg.UserConfig.DatabaseURL); err != nil {
		log.Fatal(err)
	}

	database, err := db.ConnectPostgres(cfg.UserConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	router, err := app.SetupRouter(cfg, database.DB)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server running on port %s", cfg.UserConfig.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.UserConfig.PORT, router))
}
