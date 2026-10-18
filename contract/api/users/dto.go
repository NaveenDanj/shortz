package users

type CreateUserRequest struct {
	Username    string   `validate:"required,username" json:"username"`
	DisplayName string   `validate:"required,max=50" json:"displayName"`
	Bio         string   `validate:"max=160" json:"bio"`
	Links       []string `validate:"max=5,dive,url" json:"links"`
	IsPrivate   bool     `json:"isPrivate"`
}

type UpdateUserRequest struct {
	DisplayName string   `validate:"required,max=50" json:"displayName"`
	Bio         string   `validate:"max=160" json:"bio"`
	Links       []string `validate:"max=5,dive,url" json:"links"`
	IsPrivate   bool     `json:"isPrivate"`
}
