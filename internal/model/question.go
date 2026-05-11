package model

type CreateQuestionRequest struct {
	QuizID         int64  `json:"quiz_id"           binding:"required"`
	Text           string `json:"text"              binding:"required"`
	Context        string `json:"context"`
	VideoAnswerUrl string `json:"video_answer_url"`
	Order          int64  `json:"order"`
}

type UpdateQuestionRequest struct {
	Text           string `json:"text"              binding:"required"`
	Context        string `json:"context"`
	VideoAnswerUrl string `json:"video_answer_url"`
	Order          int64  `json:"order"`
}

type QuestionResponse struct {
	ID             int64  `json:"id"`
	QuizID         int64  `json:"quiz_id"`
	Text           string `json:"text"`
	Context        string `json:"context"`
	VideoAnswerUrl string `json:"video_answer_url"`
	Order          int64  `json:"order"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
}

type GetAllQuestionsResponse struct {
	Questions []QuestionResponse `json:"questions"`
}

type DeleteQuestionResponse struct {
	Message string `json:"message"`
}
