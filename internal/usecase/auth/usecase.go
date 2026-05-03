package auth

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type AuthGRPCClient interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.TokenResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.TokenResponse, error)
	RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.TokenResponse, error)
}

type UseCase struct {
	authClient AuthGRPCClient
}

func New(authClient AuthGRPCClient) *UseCase {
	return &UseCase{authClient: authClient}
}

func (uc *UseCase) Register(ctx context.Context, req model.RegisterRequest) (*model.TokenResponse, error) {
	return uc.authClient.Register(ctx, req)
}

func (uc *UseCase) Login(ctx context.Context, req model.LoginRequest) (*model.TokenResponse, error) {
	return uc.authClient.Login(ctx, req)
}

func (uc *UseCase) RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.TokenResponse, error) {
	return uc.authClient.RefreshToken(ctx, req)
}
