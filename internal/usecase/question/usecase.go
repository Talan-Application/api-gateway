package question

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type QuestionGRPCClient interface {
	CreateQuestionWithAnswers(ctx context.Context, req model.CreateQuestionWithAnswersRequest) (*model.QuestionWithAnswersResponse, error)
	GetQuestion(ctx context.Context, id int64) (*model.QuestionWithAnswersResponse, error)
	GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error)
	UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionWithAnswersResponse, error)
	DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error)
}

type UseCase struct {
	questionClient QuestionGRPCClient
}

func New(questionClient QuestionGRPCClient) *UseCase {
	return &UseCase{questionClient: questionClient}
}

func (uc *UseCase) CreateQuestionWithAnswers(ctx context.Context, req model.CreateQuestionWithAnswersRequest) (*model.QuestionWithAnswersResponse, error) {
	return uc.questionClient.CreateQuestionWithAnswers(ctx, req)
}

func (uc *UseCase) GetQuestion(ctx context.Context, id int64) (*model.QuestionWithAnswersResponse, error) {
	return uc.questionClient.GetQuestion(ctx, id)
}

func (uc *UseCase) GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error) {
	return uc.questionClient.GetAllQuestions(ctx, quizID, limit, offset)
}

func (uc *UseCase) UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionWithAnswersResponse, error) {
	return uc.questionClient.UpdateQuestion(ctx, id, req)
}

func (uc *UseCase) DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error) {
	return uc.questionClient.DeleteQuestion(ctx, id)
}
