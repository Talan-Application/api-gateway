package quizresult

import (
	"context"

	"google.golang.org/grpc"

	quizv1 "github.com/Talan-Application/proto-generation/quiz/v1"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type Client struct {
	stub quizv1.QuizResultServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: quizv1.NewQuizResultServiceClient(conn)}
}

func (c *Client) SubmitQuiz(ctx context.Context, quizID int64, req model.SubmitQuizRequest) (*model.SubmitQuizResponse, error) {
	answers := make([]*quizv1.AnswerSubmission, len(req.Answers))
	for i, a := range req.Answers {
		answers[i] = &quizv1.AnswerSubmission{
			QuestionId: a.QuestionID,
			AnswerId:   a.AnswerID,
		}
	}

	resp, err := c.stub.SubmitQuiz(ctx, &quizv1.SubmitQuizRequest{
		QuizId:  quizID,
		Answers: answers,
	})
	if err != nil {
		return nil, err
	}

	results := make([]model.QuestionResult, len(resp.GetResults()))
	for i, r := range resp.GetResults() {
		results[i] = model.QuestionResult{
			QuestionID:       r.GetQuestionId(),
			SelectedAnswerID: r.GetSelectedAnswerId(),
			CorrectAnswerID:  r.GetCorrectAnswerId(),
			IsCorrect:        r.GetIsCorrect(),
		}
	}

	return &model.SubmitQuizResponse{
		ResultID:       resp.GetResultId(),
		TotalQuestions: resp.GetTotalQuestions(),
		CorrectAnswers: resp.GetCorrectAnswers(),
		Score:          resp.GetScore(),
		Results:        results,
	}, nil
}

func (c *Client) GetQuizResults(ctx context.Context, quizID, userID int64) (*model.GetQuizResultsResponse, error) {
	resp, err := c.stub.GetQuizResults(ctx, &quizv1.GetQuizResultsRequest{
		QuizId: quizID,
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}

	summaries := make([]model.QuizResultSummary, len(resp.GetResults()))
	for i, r := range resp.GetResults() {
		summaries[i] = model.QuizResultSummary{
			ID:             r.GetId(),
			QuizID:         r.GetQuizId(),
			UserID:         r.GetUserId(),
			Score:          r.GetScore(),
			TotalQuestions: r.GetTotalQuestions(),
			CorrectAnswers: r.GetCorrectAnswers(),
			SubmittedAt:    r.GetSubmittedAt(),
		}
	}

	return &model.GetQuizResultsResponse{Results: summaries}, nil
}
