package config

type Config struct {
	AuthConfig AuthConfig
	UserConfig UserConfig
}

type AuthConfig struct {
	DatabaseURL    string
	DatabaseDriver string
	Secret         string
	BaseURL        string
}
type UserConfig struct {
	DatabaseURL    string
	DatabaseDriver string
	PORT           string
	AuthIssuer     string // must equal the auth service's AUTH_BASE_URL
	AuthJWKSURL    string
}
