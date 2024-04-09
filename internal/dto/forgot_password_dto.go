package dto

type ForgotPasswordPasswordRequest struct {
	Email string `json:"email"`
}

type ForgotPasswordPasswordResponse struct {
	Token string `json:"token"`
}
