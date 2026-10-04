package auth

import (
	"net/http"

	authcontract "contract/api/auth"

	"github.com/go-dev-auth/go-dev-auth/storage"
)

func Me(w http.ResponseWriter, r *http.Request) {
	u, ok := UserFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"code":    "UNAUTHORIZED",
			"message": "Authentication required",
		})
		return
	}
	writeJSON(w, http.StatusOK, toContractUser(u))
}

func toContractUser(u *storage.User) authcontract.User {
	return authcontract.User{
		ID:            u.ID,
		Name:          u.Name,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		Image:         u.Image,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
	}
}
