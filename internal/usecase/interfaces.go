package usecase

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type Auth interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.MessageResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.MessageResponse, error)
	VerifyEmail(ctx context.Context, req model.VerifyCodeRequest) (*model.TokenResponse, error)
	VerifyLoginCode(ctx context.Context, req model.VerifyCodeRequest) (*model.TokenResponse, error)
	RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.TokenResponse, error)
}
