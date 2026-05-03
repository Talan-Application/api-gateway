package model

type RegisterRequest struct {
	Email      string  `json:"email"       binding:"required,email"`
	Password   string  `json:"password"    binding:"required,min=8"`
	FirstName  string  `json:"first_name"  binding:"required"`
	LastName   string  `json:"last_name"   binding:"required"`
	MiddleName *string `json:"middle_name"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
