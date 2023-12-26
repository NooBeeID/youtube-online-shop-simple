package auth

type RegisterRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
