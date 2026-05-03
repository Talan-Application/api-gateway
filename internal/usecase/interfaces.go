package usecase

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type Auth interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.TokenResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.TokenResponse, error)
	RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.TokenResponse, error)
}
