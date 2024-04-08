package dto

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordResponse struct {
	Token string `json:"token"`
}
