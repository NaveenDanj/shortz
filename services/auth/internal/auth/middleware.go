package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	godevauth "github.com/go-dev-auth/go-dev-auth"
	jwtplugin "github.com/go-dev-auth/go-dev-auth/plugins/jwt"
	"github.com/go-dev-auth/go-dev-auth/storage"
)

type ctxKey struct{}

func SessionFromContext(ctx context.Context) (*godevauth.SessionData, bool) {
	sd, ok := ctx.Value(ctxKey{}).(*godevauth.SessionData)
	return sd, ok
}

func UserFromContext(ctx context.Context) (*storage.User, bool) {
	sd, ok := SessionFromContext(ctx)
	if !ok {
		return nil, false
	}
	return sd.User, true
}

func RequireAuth(a *godevauth.Auth) func(http.Handler) http.Handler {
	jwtp, _ := a.Plugin("jwt").(*jwtplugin.Plugin)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				sd  *godevauth.SessionData
				err error
			)

			if token, ok := bearerToken(r); ok {
				sd, err = sessionFromJWT(r.Context(), a, jwtp, token)
			} else {
				sd, err = a.GetSession(r)
			}

			if err != nil || sd == nil || sd.User == nil {
				writeJSON(w, http.StatusUnauthorized, map[string]string{
					"code":    "UNAUTHORIZED",
					"message": "Authentication required",
				})
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKey{}, sd)))
		})
	}
}

func bearerToken(r *http.Request) (string, bool) {
	token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	return token, ok && token != ""
}

func sessionFromJWT(ctx context.Context, a *godevauth.Auth, p *jwtplugin.Plugin, token string) (*godevauth.SessionData, error) {
	if p == nil {
		return nil, errors.New("jwt plugin not enabled")
	}

	claims, err := p.Verify(ctx, token)
	if err != nil {
		return nil, err
	}

	base := a.Config().BaseURL
	if iss, _ := claims["iss"].(string); iss != base {
		return nil, errors.New("jwt: unexpected issuer")
	}

	if !audienceMatches(claims["aud"], base) {
		return nil, errors.New("jwt: unexpected audience")
	}

	sub, _ := claims["sub"].(string)
	if sub == "" {
		return nil, errors.New("jwt: missing subject")
	}

	user, err := a.FindUserByID(ctx, sub)
	if err != nil {
		return nil, err
	}

	return &godevauth.SessionData{User: user}, nil
}

func audienceMatches(aud any, want string) bool {
	switch v := aud.(type) {
	case string:
		return v == want
	case []any:
		for _, e := range v {
			if s, _ := e.(string); s == want {
				return true
			}
		}
	}
	return false
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
