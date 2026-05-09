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

type Quiz interface {
	CreateQuiz(ctx context.Context, req model.CreateQuizRequest) (*model.QuizResponse, error)
	GetQuiz(ctx context.Context, id int64) (*model.QuizResponse, error)
	GetAllQuizzes(ctx context.Context, limit, offset *int32) (*model.GetAllQuizzesResponse, error)
	UpdateQuiz(ctx context.Context, id int64, req model.UpdateQuizRequest) (*model.QuizResponse, error)
	DeleteQuiz(ctx context.Context, id int64) (*model.DeleteQuizResponse, error)
}
