package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	return &Config{
		AuthConfig: AuthConfig{
			DatabaseURL:    loadConfigFromEnv("AUTH_DB_URL", ""),
			DatabaseDriver: loadConfigFromEnv("AUTH_DB_DRIVER", "pgx"),
			Secret:         loadConfigFromEnv("AUTH_SECRET", ""),
			BaseURL:        loadConfigFromEnv("AUTH_BASE_URL", "http://localhost:8080"),
		},
	}, nil
}

func loadConfigFromEnv(key string, defaultString string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultString
}
