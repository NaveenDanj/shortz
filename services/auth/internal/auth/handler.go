package auth

import (
	"net/http"

	apiContract "contract/api"
	authcontract "contract/api/auth"

	"github.com/go-dev-auth/go-dev-auth/storage"
)

func Me(w http.ResponseWriter, r *http.Request) {
	u, ok := UserFromContext(r.Context())
	if !ok {
		apiContract.ResponseWriter(w, http.StatusUnauthorized, "Authentication required")
		return
	}
	apiContract.ResponseWithDataWriter(w, http.StatusOK, toContractUser(u))
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
