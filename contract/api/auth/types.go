package auth

import "time"

type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"emailVerified"`
	Image         string    `json:"image,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Session struct {
	ID        string         `json:"id"`
	UserID    string         `json:"userId"`
	Token     string         `json:"token"`
	ExpiresAt time.Time      `json:"expiresAt"`
	IPAddress string         `json:"ipAddress,omitempty"`
	UserAgent string         `json:"userAgent,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Extra     map[string]any `json:"-"`
}

type Account struct {
	ID                    string         `json:"id"`
	UserID                string         `json:"userId"`
	AccountID             string         `json:"accountId"`
	ProviderID            string         `json:"providerId"`
	AccessToken           string         `json:"-"`
	RefreshToken          string         `json:"-"`
	IDToken               string         `json:"-"`
	AccessTokenExpiresAt  time.Time      `json:"accessTokenExpiresAt,omitempty"`
	RefreshTokenExpiresAt time.Time      `json:"refreshTokenExpiresAt,omitempty"`
	Scope                 string         `json:"scope,omitempty"`
	Password              string         `json:"-"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	Extra                 map[string]any `json:"-"`
}

type Verification struct {
	ID         string    `json:"id"`
	Identifier string    `json:"identifier"`
	Value      string    `json:"value"`
	ExpiresAt  time.Time `json:"expiresAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
