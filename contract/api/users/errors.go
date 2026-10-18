package users

import "errors"

var (
	ErrNotFound      = errors.New("[User] user not found")
	ErrProfileExists = errors.New("[User] profile already exists")
	ErrUsernameTaken = errors.New("[User] username already taken")
)
