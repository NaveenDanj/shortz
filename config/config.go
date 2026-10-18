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

	authIssuer := loadConfigFromEnv("USERS_AUTH_ISSUER", loadConfigFromEnv("AUTH_BASE_URL", ""))
	authJWKSURL := loadConfigFromEnv("USERS_AUTH_JWKS_URL", "")

	return &Config{
		AuthConfig: AuthConfig{
			DatabaseURL:    loadConfigFromEnv("AUTH_DB_URL", ""),
			DatabaseDriver: loadConfigFromEnv("AUTH_DB_DRIVER", "pgx"),
			Secret:         loadConfigFromEnv("AUTH_SECRET", ""),
			BaseURL:        loadConfigFromEnv("AUTH_BASE_URL", ""),
		},
		UserConfig: UserConfig{
			DatabaseURL:    loadConfigFromEnv("USERS_DB_URL", ""),
			DatabaseDriver: loadConfigFromEnv("USERS_DB_DRIVER", "pgx"),
			PORT:           loadConfigFromEnv("USERS_PORT", "3002"),
			AuthIssuer:     authIssuer,
			AuthJWKSURL:    authJWKSURL,
		},
	}, nil
}

func loadConfigFromEnv(key string, defaultString string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultString
}
