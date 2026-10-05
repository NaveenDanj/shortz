package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"authclient"

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

			if token, ok := authclient.BearerToken(r); ok {
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

func sessionFromJWT(ctx context.Context, a *godevauth.Auth, p *jwtplugin.Plugin, token string) (*godevauth.SessionData, error) {
	if p == nil {
		return nil, errors.New("jwt plugin not enabled")
	}

	claims, err := p.Verify(ctx, token)
	if err != nil {
		return nil, err
	}

	base := a.Config().BaseURL
	parsed, err := authclient.ParseClaims(claims, base, base)
	if err != nil {
		return nil, err
	}
	user, err := a.FindUserByID(ctx, parsed.Subject)
	if err != nil {
		return nil, err
	}

	return &godevauth.SessionData{User: user}, nil
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
