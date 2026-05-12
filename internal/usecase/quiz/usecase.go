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

type QuestionGRPCClient interface {
	GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error)
}

type AnswerGRPCClient interface {
	GetAllAnswers(ctx context.Context, questionID int64, limit, offset *int32) (*model.GetAllAnswersResponse, error)
}

type QuizResultGRPCClient interface {
	SubmitQuiz(ctx context.Context, quizID int64, req model.SubmitQuizRequest) (*model.SubmitQuizResponse, error)
	GetQuizResults(ctx context.Context, quizID, userID int64) (*model.GetQuizResultsResponse, error)
}

type UseCase struct {
	quizClient       QuizGRPCClient
	questionClient   QuestionGRPCClient
	answerClient     AnswerGRPCClient
	resultClient     QuizResultGRPCClient
}

func New(quizClient QuizGRPCClient, questionClient QuestionGRPCClient, answerClient AnswerGRPCClient, resultClient QuizResultGRPCClient) *UseCase {
	return &UseCase{
		quizClient:     quizClient,
		questionClient: questionClient,
		answerClient:   answerClient,
		resultClient:   resultClient,
	}
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

func (uc *UseCase) TakeQuiz(ctx context.Context, id int64) (*model.TakeQuizResponse, error) {
	quiz, err := uc.quizClient.GetQuiz(ctx, id)
	if err != nil {
		return nil, err
	}

	questionsResp, err := uc.questionClient.GetAllQuestions(ctx, id, nil, nil)
	if err != nil {
		return nil, err
	}

	questions := make([]model.TakeQuizQuestion, len(questionsResp.Questions))
	for i, q := range questionsResp.Questions {
		answersResp, err := uc.answerClient.GetAllAnswers(ctx, q.ID, nil, nil)
		if err != nil {
			return nil, err
		}

		answers := make([]model.TakeQuizAnswer, len(answersResp.Answers))
		for j, a := range answersResp.Answers {
			answers[j] = model.TakeQuizAnswer{ID: a.ID, Text: a.Text}
		}

		questions[i] = model.TakeQuizQuestion{
			ID:             q.ID,
			Text:           q.Text,
			Context:        q.Context,
			VideoAnswerUrl: q.VideoAnswerUrl,
			Order:          q.Order,
			Answers:        answers,
		}
	}

	return &model.TakeQuizResponse{Quiz: *quiz, Questions: questions}, nil
}

func (uc *UseCase) SubmitQuiz(ctx context.Context, id int64, req model.SubmitQuizRequest) (*model.SubmitQuizResponse, error) {
	return uc.resultClient.SubmitQuiz(ctx, id, req)
}

func (uc *UseCase) GetQuizResults(ctx context.Context, quizID, userID int64) (*model.GetQuizResultsResponse, error) {
	return uc.resultClient.GetQuizResults(ctx, quizID, userID)
}
