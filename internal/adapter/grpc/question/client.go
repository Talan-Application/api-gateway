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

func (c *Client) CreateQuestionWithAnswers(ctx context.Context, req model.CreateQuestionWithAnswersRequest) (*model.QuestionWithAnswersResponse, error) {
	inputs := make([]*quizv1.AnswerInput, len(req.Answers))
	for i, a := range req.Answers {
		inputs[i] = &quizv1.AnswerInput{Text: a.Text, Correct: a.Correct}
	}

	resp, err := c.stub.CreateQuestionWithAnswers(ctx, &quizv1.CreateQuestionWithAnswersRequest{
		QuizId:         req.QuizID,
		Text:           req.Text,
		Context:        req.Context,
		VideoAnswerUrl: req.VideoAnswerUrl,
		Order:          req.Order,
		Answers:        inputs,
	})
	if err != nil {
		return nil, err
	}

	return protoToQuestionWithAnswers(resp), nil
}

func (c *Client) GetQuestion(ctx context.Context, id int64) (*model.QuestionWithAnswersResponse, error) {
	resp, err := c.stub.GetQuestion(ctx, &quizv1.GetQuestionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return protoToQuestionWithAnswers(resp), nil
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

	questions := make([]model.QuestionWithAnswersResponse, len(resp.GetQuestions()))
	for i, q := range resp.GetQuestions() {
		questions[i] = *protoToQuestionWithAnswers(q)
	}
	return &model.GetAllQuestionsResponse{Questions: questions}, nil
}

func (c *Client) UpdateQuestion(ctx context.Context, id int64, req model.UpdateQuestionRequest) (*model.QuestionWithAnswersResponse, error) {
	inputs := make([]*quizv1.AnswerInput, len(req.Answers))
	for i, a := range req.Answers {
		inputs[i] = &quizv1.AnswerInput{Text: a.Text, Correct: a.Correct}
	}

	resp, err := c.stub.UpdateQuestionWithAnswers(ctx, &quizv1.UpdateQuestionWithAnswersRequest{
		Id:             id,
		Text:           req.Text,
		Context:        req.Context,
		VideoAnswerUrl: req.VideoAnswerUrl,
		Order:          req.Order,
		Answers:        inputs,
	})
	if err != nil {
		return nil, err
	}
	return protoToQuestionWithAnswers(resp), nil
}

func (c *Client) DeleteQuestion(ctx context.Context, id int64) (*model.DeleteQuestionResponse, error) {
	resp, err := c.stub.DeleteQuestion(ctx, &quizv1.DeleteQuestionRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.DeleteQuestionResponse{Message: resp.GetMessage()}, nil
}

func protoToQuestionWithAnswers(q *quizv1.QuestionWithAnswersResponse) *model.QuestionWithAnswersResponse {
	answers := make([]model.AnswerResponse, len(q.GetAnswers()))
	for i, a := range q.GetAnswers() {
		answers[i] = model.AnswerResponse{
			ID:         a.GetId(),
			QuestionID: a.GetQuestionId(),
			Text:       a.GetText(),
			Correct:    a.GetCorrect(),
			CreatedAt:  a.GetCreatedAt(),
			UpdatedAt:  a.GetUpdatedAt(),
		}
	}
	return &model.QuestionWithAnswersResponse{
		ID:             q.GetId(),
		QuizID:         q.GetQuizId(),
		Text:           q.GetText(),
		Context:        q.GetContext(),
		VideoAnswerUrl: q.GetVideoAnswerUrl(),
		Order:          q.GetOrder(),
		CreatedAt:      q.GetCreatedAt(),
		UpdatedAt:      q.GetUpdatedAt(),
		Answers:        answers,
	}
}
