package config

type Config struct {
	AuthConfig AuthConfig
}

type AuthConfig struct {
	DatabaseURL    string
	DatabaseDriver string
	Secret         string
	BaseURL        string
}
