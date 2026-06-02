package common_subject

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	commonsubjectv1 "github.com/Talan-Application/proto-generation/common_subject/v1"

	"github.com/Talan-Application/api-gateway/internal/ctxkeys"
	"github.com/Talan-Application/api-gateway/internal/model"
)

type Client struct {
	stub commonsubjectv1.CommonSubjectServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{stub: commonsubjectv1.NewCommonSubjectServiceClient(conn)}
}

func (c *Client) CreateCommonSubject(ctx context.Context, req model.CreateCommonSubjectRequest) (*model.CommonSubjectResponse, error) {
	resp, err := c.stub.CreateCommonSubject(ctx, &commonsubjectv1.CreateCommonSubjectRequest{
		Translations: req.Translations,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetCommonSubject(ctx context.Context, id int64) (*model.CommonSubjectResponse, error) {
	ctx = withLocale(ctx)
	resp, err := c.stub.GetCommonSubject(ctx, &commonsubjectv1.GetCommonSubjectRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) GetAllCommonSubjects(ctx context.Context, limit, offset *int32) (*model.GetAllCommonSubjectsResponse, error) {
	resp, err := c.stub.GetAllCommonSubjects(ctx, &commonsubjectv1.GetAllCommonSubjectsRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	subjects := make([]model.CommonSubjectListItem, len(resp.GetCommonSubjects()))
	for i, s := range resp.GetCommonSubjects() {
		subjects[i] = model.CommonSubjectListItem{
			ID:           s.GetId(),
			CreatedAt:    s.GetCreatedAt(),
			UpdatedAt:    s.GetUpdatedAt(),
			Translations: s.GetTranslations(),
		}
	}
	return &model.GetAllCommonSubjectsResponse{CommonSubjects: subjects}, nil
}

func (c *Client) GetCommonSubjectsLookup(ctx context.Context) (*model.GetCommonSubjectsLookupResponse, error) {
	ctx = withLocale(ctx)
	resp, err := c.stub.GetCommonSubjectsLookup(ctx, &commonsubjectv1.GetCommonSubjectsLookupRequest{})
	if err != nil {
		return nil, err
	}

	items := make([]model.CommonSubjectLookupItem, len(resp.GetCommonSubjects()))
	for i, s := range resp.GetCommonSubjects() {
		items[i] = model.CommonSubjectLookupItem{
			ID:   s.GetId(),
			Name: s.GetName(),
		}
	}
	return &model.GetCommonSubjectsLookupResponse{CommonSubjects: items}, nil
}

func (c *Client) UpdateCommonSubject(ctx context.Context, id int64, req model.UpdateCommonSubjectRequest) (*model.CommonSubjectResponse, error) {
	resp, err := c.stub.UpdateCommonSubject(ctx, &commonsubjectv1.UpdateCommonSubjectRequest{
		Id:           id,
		Translations: req.Translations,
	})
	if err != nil {
		return nil, err
	}
	return toModel(resp), nil
}

func (c *Client) DeleteCommonSubject(ctx context.Context, id int64) (*model.DeleteCommonSubjectResponse, error) {
	resp, err := c.stub.DeleteCommonSubject(ctx, &commonsubjectv1.DeleteCommonSubjectRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &model.DeleteCommonSubjectResponse{Message: resp.GetMessage()}, nil
}

func toModel(s *commonsubjectv1.CommonSubjectResponse) *model.CommonSubjectResponse {
	return &model.CommonSubjectResponse{
		ID:        s.GetId(),
		Name:      s.GetName(),
		CreatedAt: s.GetCreatedAt(),
		UpdatedAt: s.GetUpdatedAt(),
	}
}

// withLocale attaches the user's preferred locale from context as gRPC metadata
// so system-handbook-service can resolve the correct translation.
func withLocale(ctx context.Context) context.Context {
	locale := "ru"
	if v, ok := ctx.Value(ctxkeys.LocaleKey).(string); ok && v != "" {
		locale = v
	}
	return metadata.AppendToOutgoingContext(ctx, "locale", locale)
}
