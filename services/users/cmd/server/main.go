package main

import (
	"config"
	"log"
	"net/http"
	"users/pkg/db"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	database, err := db.ConnectPostgres(cfg.UserConfig.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// app.RunMigrations(database.DB)
	// if cfg.SeedSampleData {
	// 	app.RunSeeders(database.DB)
	// }

	// router := app.SetupRouter(cfg, database.DB)

	log.Printf("Server running on port %s", cfg.UserConfig.PORT)
	log.Fatal(http.ListenAndServe(":"+cfg.UserConfig.PORT, router))
}
