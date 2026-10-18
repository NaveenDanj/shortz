package users

import (
	"authclient"
	"common/httpx"
	apiContract "contract/api"
	apiUsers "contract/api/users"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := authclient.UserID(r.Context())

	user, err := c.service.GetMe(r.Context(), userID)
	if err != nil {
		httpx.WriteError(w, "get me", userID, err)
		return
	}

	apiContract.ResponseWithDataWriter(w, http.StatusOK, user)
}

func (c *Controller) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := authclient.UserID(r.Context())

	req, ok := httpx.DecodeAndValidate[apiUsers.UpdateUserRequest](w, r)
	if !ok {
		return
	}

	if err := c.service.UpdateProfile(r.Context(), userID, req); err != nil {
		httpx.WriteError(w, "update profile", userID, err)
		return
	}

	apiContract.ResponseWriter(w, http.StatusOK, "Profile updated successfully")
}
