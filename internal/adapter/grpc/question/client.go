package question

import (
	"context"

	"google.golang.org/grpc"

	quizv1 "github.com/Talan-Application/proto-generation/quiz/v1"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type Client struct {
	stub quizv1.QuestionServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: quizv1.NewQuestionServiceClient(conn)}
}

func (c *Client) CreateQuestion(ctx context.Context, req model.CreateQuestionRequest) (*model.QuestionResponse, error) {
	resp, err := c.stub.CreateQuestion(ctx, &quizv1.CreateQuestionRequest{
		QuizId:         req.QuizID,
		Text:           req.Text,
		Context:        req.Context,
		VideoAnswerUrl: req.VideoAnswerUrl,
		Order:          req.Order,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetQuestion(ctx context.Context, id int64) (*model.QuestionResponse, error) {
	resp, err := c.stub.GetQuestion(ctx, &quizv1.GetQuestionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetAllQuestions(ctx context.Context, quizID int64, limit, offset *int32) (*model.GetAllQuestionsResponse, error) {
	resp, err := c.stub.GetAllQuestions(ctx, &quizv1.GetAllQuestionsRequest{
		QuizId: quizID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	questions := make([]model.QuestionResponse, len(resp.GetQuestions()))
	for i, q := range resp.GetQuestions() {
		questions[i] = *toModel(q)
	}
	return &model.GetAllQuestionsResponse{Questions: questions}, nil
}

func (c *Client) UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionResponse, error) {
	resp, err := c.stub.UpdateQuestion(ctx, &quizv1.UpdateQuestionRequest{
		Id:             id,
		Text:           req.Text,
		Context:        req.Context,
		VideoAnswerUrl: req.VideoAnswerUrl,
		Order:          req.Order,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error) {
	resp, err := c.stub.DeleteQuestion(ctx, &quizv1.DeleteQuestionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.DeleteQuestionResponse{Message: resp.GetMessage()}, nil
}

func toModel(q *quizv1.QuestionResponse) *model.QuestionResponse {
	return &model.QuestionResponse{
		ID:             q.GetId(),
		QuizID:         q.GetQuizId(),
		Text:           q.GetText(),
		Context:        q.GetContext(),
		VideoAnswerUrl: q.GetVideoAnswerUrl(),
		Order:          q.GetOrder(),
		CreatedAt:      q.GetCreatedAt(),
		UpdatedAt:      q.GetUpdatedAt(),
	}
}
