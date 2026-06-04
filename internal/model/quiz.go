package model

type CreateQuizRequest struct {
	Title           string `json:"title"      binding:"required"`
	Language        string `json:"language"   binding:"required"`
	Type            string `json:"type"       binding:"required,oneof=ent monthly_exam exam"`
	CommonSubjectID int64  `json:"common_subject_id" binding:"required"`
	IsEntStandard   bool   `json:"is_ent_standard"`
}

type UpdateQuizRequest struct {
	Title         string `json:"title"    binding:"required"`
	Language      string `json:"language" binding:"required"`
	Type          string `json:"type"     binding:"required,oneof=ent monthly_exam exam"`
	IsEntStandard bool   `json:"is_ent_standard"`
}

type QuizResponse struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Language      string `json:"language"`
	AuthorID      int64  `json:"author_id"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	CommonSubjectID int64  `json:"common_subject_id"`
	IsEntStandard bool   `json:"is_ent_standard"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
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
	QuestionID int64   `json:"question_id" binding:"required"`
	AnswerIDs  []int64 `json:"answer_ids"  binding:"required,min=1"`
}

type SubmitQuizRequest struct {
	Answers []QuizAnswerSubmission `json:"answers" binding:"required,min=1"`
}

type QuestionResult struct {
	QuestionID        int64   `json:"question_id"`
	SelectedAnswerIDs []int64 `json:"selected_answer_ids"`
	CorrectAnswerIDs  []int64 `json:"correct_answer_ids"`
	Score             float64 `json:"score"`
	MaxScore          float64 `json:"max_score"`
}

type SubmitQuizResponse struct {
	ResultID             int64            `json:"result_id"`
	TotalQuestionsCount  int32            `json:"total_questions_count"`
	CorrectAnswersCount  int32            `json:"correct_answers_count"`
	IncorrectAnswersCount int32           `json:"incorrect_answers_count"`
	UnansweredQuestions  int32            `json:"unanswered_questions"`
	Score                float64          `json:"score"`
	MaxScore             float64          `json:"max_score"`
	Results              []QuestionResult `json:"results"`
}

type QuizResultSummary struct {
	ID                   int64   `json:"id"`
	QuizID               int64   `json:"quiz_id"`
	UserID               int64   `json:"user_id"`
	Score                float64 `json:"score"`
	MaxScore             float64 `json:"max_score"`
	TotalQuestionsCount  int32   `json:"total_questions_count"`
	CorrectAnswersCount  int32   `json:"correct_answers_count"`
	IncorrectAnswersCount int32  `json:"incorrect_answers_count"`
	UnansweredQuestions  int32   `json:"unanswered_questions"`
	SubmittedAt          int64   `json:"submitted_at"`
}

type GetQuizResultsResponse struct {
	Results []QuizResultSummary `json:"results"`
}
