package auth

type Constants struct {
	MIN_PASSWORD_LENGTH int
	MAX_PASSWORD_LENGTH int
}

var Auth_Constants = Constants{
	MIN_PASSWORD_LENGTH: 8,
	MAX_PASSWORD_LENGTH: 64,
}
