package model

type CreateQuizRequest struct {
	Title     string `json:"title"      binding:"required"`
	Language  string `json:"language"   binding:"required"`
	Type      string `json:"type"       binding:"required,oneof=ent monthly_exam exam"`
	SubjectID int64  `json:"subject_id" binding:"required"`
}

type UpdateQuizRequest struct {
	Title    string `json:"title"    binding:"required"`
	Language string `json:"language" binding:"required"`
	Type     string `json:"type"     binding:"required,oneof=ent monthly_exam exam"`
}

type QuizResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Language  string `json:"language"`
	AuthorID  int64  `json:"author_id"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	SubjectID int64  `json:"subject_id"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type GetAllQuizzesResponse struct {
	Quizzes []QuizResponse `json:"quizzes"`
}

type DeleteQuizResponse struct {
	Message string `json:"message"`
}

type TakeQuizAnswer struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type TakeQuizQuestion struct {
	ID             int64            `json:"id"`
	Text           string           `json:"text"`
	Context        string           `json:"context"`
	VideoAnswerUrl string           `json:"video_answer_url"`
	Order          int64            `json:"order"`
	Answers        []TakeQuizAnswer `json:"answers"`
}

type TakeQuizResponse struct {
	Quiz      QuizResponse       `json:"quiz"`
	Questions []TakeQuizQuestion `json:"questions"`
}

type QuizAnswerSubmission struct {
	QuestionID int64 `json:"question_id" binding:"required"`
	AnswerID   int64 `json:"answer_id"   binding:"required"`
}

type SubmitQuizRequest struct {
	Answers []QuizAnswerSubmission `json:"answers" binding:"required,min=1"`
}

type QuestionResult struct {
	QuestionID       int64 `json:"question_id"`
	SelectedAnswerID int64 `json:"selected_answer_id"`
	CorrectAnswerID  int64 `json:"correct_answer_id"`
	IsCorrect        bool  `json:"is_correct"`
}

type SubmitQuizResponse struct {
	TotalQuestions int              `json:"total_questions"`
	CorrectAnswers int              `json:"correct_answers"`
	Score          float64          `json:"score"`
	Results        []QuestionResult `json:"results"`
}
