package authclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-dev-auth/go-dev-auth/crypto"
)

var ErrInvalidToken = errors.New("authclient: invalid token")

const (
	defaultRefreshInterval = 10 * time.Minute
	minRefetchInterval     = 30 * time.Second
)

func NewVerifier(cfg Config) (*Verifier, error) {

	if cfg.JWKSURL == "" || cfg.Issuer == "" || cfg.Audience == "" {
		return nil, errors.New("authclient: JWKSURL, Issuer and Audience are required")
	}

	if cfg.RefreshInterval == 0 {
		cfg.RefreshInterval = defaultRefreshInterval
	}

	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 5 * time.Second}
	}

	return &Verifier{cfg: cfg}, nil
}

func (v *Verifier) Verify(ctx context.Context, token string) (*Claims, error) {

	kid, err := tokenKeyID(token)
	if err != nil {
		return nil, ErrInvalidToken
	}

	key, err := v.key(ctx, kid)
	if err != nil {
		return nil, err
	}

	claims, err := crypto.VerifyJWT(key, token)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	return ParseClaims(claims, v.cfg.Issuer, v.cfg.Audience)
}

func ParseClaims(claims crypto.Claims, issuer, audience string) (*Claims, error) {
	if iss, _ := claims["iss"].(string); iss != issuer {
		return nil, fmt.Errorf("%w: unexpected issuer", ErrInvalidToken)
	}

	if !audienceMatches(claims["aud"], audience) {
		return nil, fmt.Errorf("%w: unexpected audience", ErrInvalidToken)
	}

	sub, _ := claims["sub"].(string)
	if sub == "" {
		return nil, fmt.Errorf("%w: missing subject", ErrInvalidToken)
	}

	out := &Claims{Subject: sub, Raw: claims}
	out.Email, _ = claims["email"].(string)
	out.Name, _ = claims["name"].(string)
	out.Image, _ = claims["image"].(string)
	out.EmailVerified, _ = claims["emailVerified"].(bool)
	return out, nil
}

func (v *Verifier) key(ctx context.Context, kid string) (any, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	stale := v.keys == nil || time.Since(v.fetchedAt) > v.cfg.RefreshInterval
	if key, ok := v.keys[kid]; ok && !stale {
		return key, nil
	}

	if v.keys == nil || time.Since(v.fetchedAt) >= minRefetchInterval {
		if err := v.fetch(ctx); err != nil {
			if key, ok := v.keys[kid]; ok {
				return key, nil
			}
			return nil, err
		}
	}
	if key, ok := v.keys[kid]; ok {
		return key, nil
	}
	return nil, fmt.Errorf("%w: unknown key id", ErrInvalidToken)
}

func (v *Verifier) fetch(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, v.cfg.JWKSURL, nil)
	if err != nil {
		return err
	}
	resp, err := v.cfg.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("authclient: fetching jwks: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("authclient: jwks endpoint returned %d", resp.StatusCode)
	}
	var set crypto.JWKS
	if err := json.NewDecoder(resp.Body).Decode(&set); err != nil {
		return fmt.Errorf("authclient: decoding jwks: %w", err)
	}
	keys := make(map[string]any, len(set.Keys))
	for _, jwk := range set.Keys {
		if pub, err := jwk.PublicKey(); err == nil {
			keys[jwk.Kid] = pub
		}
	}
	v.keys = keys
	v.fetchedAt = time.Now()
	return nil
}

func tokenKeyID(token string) (string, error) {
	header, _, ok := strings.Cut(token, ".")
	if !ok {
		return "", ErrInvalidToken
	}
	raw, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		return "", err
	}
	var h crypto.JWTHeader
	if err := json.Unmarshal(raw, &h); err != nil {
		return "", err
	}
	return h.Kid, nil
}

func audienceMatches(aud any, want string) bool {
	switch a := aud.(type) {
	case string:
		return a == want
	case []any:
		for _, e := range a {
			if s, _ := e.(string); s == want {
				return true
			}
		}
	}
	return false
}
