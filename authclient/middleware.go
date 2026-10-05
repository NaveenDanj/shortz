package authclient

import (
	"context"
	apiContract "contract/api"
	"net/http"
	"strings"
)

type ctxKey struct{}

func ClaimsFromContext(ctx context.Context) (*Claims, bool) {
	c, ok := ctx.Value(ctxKey{}).(*Claims)
	return c, ok
}

func (v *Verifier) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := BearerToken(r)

		if !ok {
			apiContract.ResponseWriter(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims, err := v.Verify(r.Context(), token)
		if err != nil {
			apiContract.ResponseWriter(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKey{}, claims)))
	})
}

func BearerToken(r *http.Request) (string, bool) {
	token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	return token, ok && token != ""
}
