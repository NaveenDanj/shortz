package betterauth

import (
	"config"
	"database/sql"
	"time"

	authcontract "contract/api/auth"

	godevauth "github.com/go-dev-auth/go-dev-auth"
	jwtplugin "github.com/go-dev-auth/go-dev-auth/plugins/jwt"
	"github.com/go-dev-auth/go-dev-auth/storage/sqlstore"
)

func New(cfg config.AuthConfig, db *sql.DB) (*godevauth.Auth, error) {
	store := sqlstore.New(db, sqlstore.Postgres, nil)

	auth, err := godevauth.New(godevauth.Config{
		BaseURL:  cfg.BaseURL,
		Secret:   cfg.Secret,
		Database: store,

		EmailAndPassword: godevauth.EmailPasswordConfig{
			Enabled:           true,
			MinPasswordLength: authcontract.Auth_Constants.MIN_PASSWORD_LENGTH,
			MaxPasswordLength: authcontract.Auth_Constants.MAX_PASSWORD_LENGTH,
		},

		// Short-lived EdDSA JWTs minted from a session (GET /api/auth/token)
		// and verifiable offline via GET /api/auth/jwks. Sessions remain the
		// source of truth; a JWT cannot be revoked, only expire.
		Plugins: []godevauth.Plugin{
			jwtplugin.New(jwtplugin.Options{ExpiresIn: 15 * time.Minute}),
		},
	})

	if err != nil {
		return nil, err
	}

	store.SetSchema(auth.Schema())
	return auth, nil
}
