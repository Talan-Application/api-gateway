package model

type CreateAnswerInput struct {
	Text    string `json:"text"    binding:"required"`
	Correct bool   `json:"correct"`
}

type AnswerResponse struct {
	ID         int64  `json:"id"`
	QuestionID int64  `json:"question_id"`
	Text       string `json:"text"`
	Correct    bool   `json:"correct"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
