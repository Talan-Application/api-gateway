package auth

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Talan-Application/api-gateway/internal/model"
	authv1 "github.com/Talan-Application/proto-generation/auth/v1"
)

type Client struct {
	stub authv1.AuthServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: authv1.NewAuthServiceClient(conn)}
}

func (c *Client) Register(ctx context.Context, req model.RegisterRequest) (*model.MessageResponse, error) {
	resp, err := c.stub.Register(ctx, &authv1.RegisterRequest{
		Email:      req.Email,
		Password:   req.Password,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
	})
	if err != nil {
		return nil, err
	}
	return &model.MessageResponse{Message: resp.GetMessage()}, nil
}

func (c *Client) Login(ctx context.Context, req model.LoginRequest) (*model.MessageResponse, error) {
	resp, err := c.stub.Login(ctx, &authv1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.MessageResponse{Message: resp.GetMessage()}, nil
}

func (c *Client) VerifyEmail(ctx context.Context, req model.VerifyCodeRequest) (*model.TokenResponse, error) {
	resp, err := c.stub.VerifyEmail(ctx, &authv1.VerifyCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &model.TokenResponse{
		AccessToken:  resp.GetAccessToken(),
		RefreshToken: resp.GetRefreshToken(),
	}, nil
}

func (c *Client) VerifyLoginCode(ctx context.Context, req model.VerifyCodeRequest) (*model.TokenResponse, error) {
	resp, err := c.stub.VerifyLoginCode(ctx, &authv1.VerifyCodeRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &model.TokenResponse{
		AccessToken:  resp.GetAccessToken(),
		RefreshToken: resp.GetRefreshToken(),
	}, nil
}

func (c *Client) RefreshToken(ctx context.Context, req model.RefreshTokenRequest) (*model.TokenResponse, error) {
	resp, err := c.stub.RefreshToken(ctx, &authv1.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, err
	}
	return &model.TokenResponse{
		AccessToken:  resp.GetAccessToken(),
		RefreshToken: resp.GetRefreshToken(),
	}, nil
}
