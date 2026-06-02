package common_subject

import (
	"context"

	"github.com/Talan-Application/api-gateway/internal/model"
)

type CommonSubjectGRPCClient interface {
	CreateCommonSubject(ctx context.Context, req model.CreateCommonSubjectRequest) (*model.CommonSubjectResponse, error)
	GetCommonSubject(ctx context.Context, id int64) (*model.CommonSubjectResponse, error)
	GetAllCommonSubjects(ctx context.Context, limit, offset *int32) (*model.GetAllCommonSubjectsResponse, error)
	GetCommonSubjectsLookup(ctx context.Context) (*model.GetCommonSubjectsLookupResponse, error)
	UpdateCommonSubject(ctx context.Context, id int64, req model.UpdateCommonSubjectRequest) (*model.CommonSubjectResponse, error)
	DeleteCommonSubject(ctx context.Context, id int64) (*model.DeleteCommonSubjectResponse, error)
}

type UseCase struct {
	client CommonSubjectGRPCClient
}

func New(client CommonSubjectGRPCClient) *UseCase {
	return &UseCase{client: client}
}

func (uc *UseCase) CreateCommonSubject(ctx context.Context, req model.CreateCommonSubjectRequest) (*model.CommonSubjectResponse, error) {
	return uc.client.CreateCommonSubject(ctx, req)
}

func (uc *UseCase) GetCommonSubject(ctx context.Context, id int64) (*model.CommonSubjectResponse, error) {
	return uc.client.GetCommonSubject(ctx, id)
}

func (uc *UseCase) GetAllCommonSubjects(ctx context.Context, limit, offset *int32) (*model.GetAllCommonSubjectsResponse, error) {
	return uc.client.GetAllCommonSubjects(ctx, limit, offset)
}

func (uc *UseCase) GetCommonSubjectsLookup(ctx context.Context) (*model.GetCommonSubjectsLookupResponse, error) {
	return uc.client.GetCommonSubjectsLookup(ctx)
}

func (uc *UseCase) UpdateCommonSubject(ctx context.Context, id int64, req model.UpdateCommonSubjectRequest) (*model.CommonSubjectResponse, error) {
	return uc.client.UpdateCommonSubject(ctx, id, req)
}

func (uc *UseCase) DeleteCommonSubject(ctx context.Context, id int64) (*model.DeleteCommonSubjectResponse, error) {
	return uc.client.DeleteCommonSubject(ctx, id)
}
