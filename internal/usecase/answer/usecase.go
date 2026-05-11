package answer

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type AnswerGRPCClient interface {
	CreateAnswer(ctx context.Context, req model.CreateAnswerRequest) (*model.AnswerResponse, error)
	GetAnswer(ctx context.Context, id int64) (*model.AnswerResponse, error)
	GetAllAnswers(ctx context.Context, questionID int64, limit, offset *int32) (*model.GetAllAnswersResponse, error)
	UpdateAnswer(ctx context.Context, id int64, req model.UpdateAnswerRequest) (*model.AnswerResponse, error)
	DeleteAnswer(ctx context.Context, id int64) (*model.DeleteAnswerResponse, error)
}

type UseCase struct {
	answerClient AnswerGRPCClient
}

func New(answerClient AnswerGRPCClient) *UseCase {
	return &UseCase{answerClient: answerClient}
}

func (uc *UseCase) CreateAnswer(ctx context.Context, req model.CreateAnswerRequest) (*model.AnswerResponse, error) {
	return uc.answerClient.CreateAnswer(ctx, req)
}

func (uc *UseCase) GetAnswer(ctx context.Context, id int64) (*model.AnswerResponse, error) {
	return uc.answerClient.GetAnswer(ctx, id)
}

func (uc *UseCase) GetAllAnswers(ctx context.Context, questionID int64, limit, offset *int32) (*model.GetAllAnswersResponse, error) {
	return uc.answerClient.GetAllAnswers(ctx, questionID, limit, offset)
}

func (uc *UseCase) UpdateAnswer(ctx context.Context, id int64, req model.UpdateAnswerRequest) (*model.AnswerResponse, error) {
	return uc.answerClient.UpdateAnswer(ctx, id, req)
}

func (uc *UseCase) DeleteAnswer(ctx context.Context, id int64) (*model.DeleteAnswerResponse, error) {
	return uc.answerClient.DeleteAnswer(ctx, id)
}
