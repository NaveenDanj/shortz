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
	AuthIssuer     string
	AuthJWKSURL    string
}

type GlobalPGPoolConsts struct {
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifeTime int
	ConnectionMaxIdleTime int
}
