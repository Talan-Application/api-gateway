package answer

import (
	"context"

	"google.golang.org/grpc"

	quizv1 "github.com/Talan-Application/proto-generation/quiz/v1"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type Client struct {
	stub quizv1.AnswerServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: quizv1.NewAnswerServiceClient(conn)}
}

func (c *Client) CreateAnswer(ctx context.Context, req model.CreateAnswerRequest) (*model.AnswerResponse, error) {
	resp, err := c.stub.CreateAnswer(ctx, &quizv1.CreateAnswerRequest{
		QuestionId: req.QuestionID,
		Text:       req.Text,
		Correct:    req.Correct,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetAnswer(ctx context.Context, id int64) (*model.AnswerResponse, error) {
	resp, err := c.stub.GetAnswer(ctx, &quizv1.GetAnswerRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetAllAnswers(ctx context.Context, questionID int64, limit, offset *int32) (*model.GetAllAnswersResponse, error) {
	resp, err := c.stub.GetAllAnswers(ctx, &quizv1.GetAllAnswersRequest{
		QuestionId: questionID,
		Limit:      limit,
		Offset:     offset,
	})
	if err != nil {
		return nil, err
	}

	answers := make([]model.AnswerResponse, len(resp.GetAnswers()))
	for i, a := range resp.GetAnswers() {
		answers[i] = *toModel(a)
	}
	return &model.GetAllAnswersResponse{Answers: answers}, nil
}

func (c *Client) UpdateAnswer(ctx context.Context, id int64, req model.UpdateAnswerRequest) (*model.AnswerResponse, error) {
	resp, err := c.stub.UpdateAnswer(ctx, &quizv1.UpdateAnswerRequest{
		Id:      id,
		Text:    req.Text,
		Correct: req.Correct,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) DeleteAnswer(ctx context.Context, id int64) (*model.DeleteAnswerResponse, error) {
	resp, err := c.stub.DeleteAnswer(ctx, &quizv1.DeleteAnswerRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.DeleteAnswerResponse{Message: resp.GetMessage()}, nil
}

func toModel(a *quizv1.AnswerResponse) *model.AnswerResponse {
	return &model.AnswerResponse{
		ID:         a.GetId(),
		QuestionID: a.GetQuestionId(),
		Text:       a.GetText(),
		Correct:    a.GetCorrect(),
		CreatedAt:  a.GetCreatedAt(),
		UpdatedAt:  a.GetUpdatedAt(),
	}
}
