package model

type CreateQuestionWithAnswersRequest struct {
	QuizID         int64               `json:"quiz_id"           binding:"required"`
	Text           string              `json:"text"              binding:"required"`
	Context        string              `json:"context"`
	VideoAnswerUrl string              `json:"video_answer_url"`
	Order          int64               `json:"order"`
	Answers        []CreateAnswerInput `json:"answers"`
}

type UpdateQuestionRequest struct {
	Text           string              `json:"text"              binding:"required"`
	Context        string              `json:"context"`
	VideoAnswerUrl string              `json:"video_answer_url"`
	Order          int64               `json:"order"`
	Answers        []CreateAnswerInput `json:"answers"`
}

type QuestionWithAnswersResponse struct {
	ID             int64            `json:"id"`
	QuizID         int64            `json:"quiz_id"`
	Text           string           `json:"text"`
	Context        string           `json:"context"`
	VideoAnswerUrl string           `json:"video_answer_url"`
	Order          int64            `json:"order"`
	CreatedAt      int64            `json:"created_at"`
	UpdatedAt      int64            `json:"updated_at"`
	Answers        []AnswerResponse `json:"answers"`
}

type GetAllQuestionsResponse struct {
	Questions []QuestionWithAnswersResponse `json:"questions"`
}

type DeleteQuestionResponse struct {
	Message string `json:"message"`
}
