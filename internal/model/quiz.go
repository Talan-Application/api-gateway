package model

type CreateQuizRequest struct {
	Title     string `json:"title"      binding:"required"`
	Language  string `json:"language"   binding:"required"`
	AuthorID  int64  `json:"author_id"  binding:"required"`
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
