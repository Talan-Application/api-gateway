package model

type CreateCommonSubjectRequest struct {
	Translations map[string]string `json:"translations" binding:"required"`
}

type UpdateCommonSubjectRequest struct {
	Translations map[string]string `json:"translations" binding:"required"`
}

type CommonSubjectResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CommonSubjectListItem struct {
	ID           int64             `json:"id"`
	CreatedAt    int64             `json:"created_at"`
	UpdatedAt    int64             `json:"updated_at"`
	Translations map[string]string `json:"translations"`
}

type GetAllCommonSubjectsResponse struct {
	CommonSubjects []CommonSubjectListItem `json:"common_subjects"`
}

type DeleteCommonSubjectResponse struct {
	Message string `json:"message"`
}
