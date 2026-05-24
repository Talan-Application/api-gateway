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

type VerifyCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code"  binding:"required,len=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfo struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name"`
	Role       string  `json:"role"`
}

type VerifyLoginResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         UserInfo `json:"user"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
