package question

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type QuestionGRPCClient interface {
	CreateQuestion(ctx context.Context, req model.CreateQuestionRequest) (*model.QuestionResponse, error)
	GetQuestion(ctx context.Context, id int64) (*model.QuestionResponse, error)
	GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionResponse, error)
	DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error)
}

type UseCase struct {
	questionClient QuestionGRPCClient
}

func New(questionClient QuestionGRPCClient) *UseCase {
	return &UseCase{questionClient: questionClient}
}

func (uc *UseCase) CreateQuestion(ctx context.Context, req model.CreateQuestionRequest) (*model.QuestionResponse, error) {
	return uc.questionClient.CreateQuestion(ctx, req)
}

func (uc *UseCase) GetQuestion(ctx context.Context, id int64) (*model.QuestionResponse, error) {
	return uc.questionClient.GetQuestion(ctx, id)
}

func (uc *UseCase) GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error) {
	return uc.questionClient.GetAllQuestions(ctx, quizID, limit, offset)
}

func (uc *UseCase) UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionResponse, error) {
	return uc.questionClient.UpdateQuestion(ctx, id, req)
}

func (uc *UseCase) DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error) {
	return uc.questionClient.DeleteQuestion(ctx, id)
}
