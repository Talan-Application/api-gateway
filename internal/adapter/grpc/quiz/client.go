package quiz

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Talan-Application/api-gateway/internal/model"
	quizv1 "github.com/Talan-Application/proto-generation/quiz/v1"
)

type Client struct {
	stub quizv1.QuizServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: quizv1.NewQuizServiceClient(conn)}
}

func (c *Client) CreateQuiz(ctx context.Context, req model.CreateQuizRequest) (*model.QuizResponse, error) {
	resp, err := c.stub.CreateQuiz(ctx, &quizv1.CreateQuizRequest{
		Title:           req.Title,
		Language:        req.Language,
		Type:            req.Type,
		CommonSubjectId: req.CommonSubjectID,
		IsEntStandard:   req.IsEntStandard,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetQuiz(ctx context.Context, id int64) (*model.QuizResponse, error) {
	resp, err := c.stub.GetQuiz(ctx, &quizv1.GetQuizRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetAllQuizzes(ctx context.Context, status *string, limit, offset *int32) (*model.GetAllQuizzesResponse, error) {
	resp, err := c.stub.GetAllQuizzes(ctx, &quizv1.GetAllQuizzesRequest{
		Status: status,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	quizzes := make([]model.QuizResponse, len(resp.GetQuizzes()))
	for i, q := range resp.GetQuizzes() {
		quizzes[i] = *toModel(q)
	}
	return &model.GetAllQuizzesResponse{Quizzes: quizzes}, nil
}

func (c *Client) PublishQuiz(ctx context.Context, id int64) error {
	_, err := c.stub.PublishQuiz(ctx, &quizv1.PublishQuizRequest{Id: id})
	return err
}

func (c *Client) GetMyQuizzes(ctx context.Context, limit, offset *int32) (*model.GetAllQuizzesResponse, error) {
	resp, err := c.stub.GetMyQuizzes(ctx, &quizv1.GetMyQuizzesRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	quizzes := make([]model.QuizResponse, len(resp.GetQuizzes()))
	for i, q := range resp.GetQuizzes() {
		quizzes[i] = *toModel(q)
	}
	return &model.GetAllQuizzesResponse{Quizzes: quizzes}, nil
}

func (c *Client) UpdateQuiz(ctx context.Context, id int64, req model.UpdateQuizRequest) (*model.QuizResponse, error) {
	resp, err := c.stub.UpdateQuiz(ctx, &quizv1.UpdateQuizRequest{
		Id:            id,
		Title:         req.Title,
		Language:      req.Language,
		Type:          req.Type,
		IsEntStandard: req.IsEntStandard,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) DeleteQuiz(ctx context.Context, id int64) (*model.DeleteQuizResponse, error) {
	resp, err := c.stub.DeleteQuiz(ctx, &quizv1.DeleteQuizRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.DeleteQuizResponse{Message: resp.GetMessage()}, nil
}

func toModel(q *quizv1.QuizResponse) *model.QuizResponse {
	return &model.QuizResponse{
		ID:            q.GetId(),
		Title:         q.GetTitle(),
		Language:      q.GetLanguage(),
		AuthorID:      q.GetAuthorId(),
		Status:        q.GetStatus(),
		Type:          q.GetType(),
		SubjectID:     q.GetCommonSubjectId(),
		IsEntStandard: q.GetIsEntStandard(),
		CreatedAt:     q.GetCreatedAt(),
		UpdatedAt:     q.GetUpdatedAt(),
	}
}
