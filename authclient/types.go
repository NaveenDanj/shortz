package authclient

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-dev-auth/go-dev-auth/crypto"
)

type Config struct {
	JWKSURL         string
	Issuer          string
	Audience        string
	RefreshInterval time.Duration
	HTTPClient      *http.Client
}

type Claims struct {
	Subject       string
	Email         string
	Name          string
	EmailVerified bool
	Image         string
	Raw           crypto.Claims
}

type Verifier struct {
	cfg       Config
	mu        sync.Mutex
	keys      map[string]any
	fetchedAt time.Time
}
