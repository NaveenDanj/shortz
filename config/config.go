package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var PGPoolConsts GlobalPGPoolConsts

func InitConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using environment variables")
	}

	PGPoolConsts = GlobalPGPoolConsts{
		MaxOpenConnections:    loadIntFromEnv("MAX_OPEN_CONNECTIONS", 25),
		MaxIdleConnections:    loadIntFromEnv("MAX_IDLE_CONNECTIONS", 10),
		ConnectionMaxLifeTime: loadIntFromEnv("CONN_MAX_LIFETIME_MINUTES", 30),
		ConnectionMaxIdleTime: loadIntFromEnv("CONN_MAX_IDLE_TIME_MINUTES", 5),
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

func loadIntFromEnv(key string, def int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil || value <= 0 {
		return def
	}
	return value
}

func loadConfigFromEnv(key string, defaultString string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultString
}
