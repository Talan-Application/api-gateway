package quiz

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type QuizGRPCClient interface {
	CreateQuiz(ctx context.Context, req model.CreateQuizRequest) (*model.QuizResponse, error)
	GetQuiz(ctx context.Context, id int64) (*model.QuizResponse, error)
	GetAllQuizzes(ctx context.Context, limit, offset *int32) (*model.GetAllQuizzesResponse, error)
	UpdateQuiz(ctx context.Context, id int64, req model.UpdateQuizRequest) (*model.QuizResponse, error)
	DeleteQuiz(ctx context.Context, id int64) (*model.DeleteQuizResponse, error)
}

type UseCase struct {
	quizClient QuizGRPCClient
}

func New(quizClient QuizGRPCClient) *UseCase {
	return &UseCase{quizClient: quizClient}
}

func (uc *UseCase) CreateQuiz(ctx context.Context, req model.CreateQuizRequest) (*model.QuizResponse, error) {
	return uc.quizClient.CreateQuiz(ctx, req)
}

func (uc *UseCase) GetQuiz(ctx context.Context, id int64) (*model.QuizResponse, error) {
	return uc.quizClient.GetQuiz(ctx, id)
}

func (uc *UseCase) GetAllQuizzes(ctx context.Context, limit, offset *int32) (*model.GetAllQuizzesResponse, error) {
	return uc.quizClient.GetAllQuizzes(ctx, limit, offset)
}

func (uc *UseCase) UpdateQuiz(ctx context.Context, id int64, req model.UpdateQuizRequest) (*model.QuizResponse, error) {
	return uc.quizClient.UpdateQuiz(ctx, id, req)
}

func (uc *UseCase) DeleteQuiz(ctx context.Context, id int64) (*model.DeleteQuizResponse, error) {
	return uc.quizClient.DeleteQuiz(ctx, id)
}
