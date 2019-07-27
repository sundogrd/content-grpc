package content

import (
	"context"
	"github.com/sundogrd/content-grpc/repositories/content"
)

type GetContentRequest struct {
	ContentID int64
}
type GetContentResponse struct {
	Content *Content
}

type ListContentsRequest struct {
	AppID       string
	TargetId    int64
	Page        int32
	PageSize    int32
	CreatorId   int64
	ParentId    int64
	State       content.ContentState
}

type ListContentsResponse struct {
	AppID    string
	Contents []*Content
	Total    int64
	Page     int32
	PageSize int32
}

type CreateContentRequest struct {
	AppID        string
	Title   string
	Description *string
	AuthorID int64
	Category *string
	Type     *content.ContentType
	Body     string
	BodyType *content.ContentBodyType
	Extra    map[string]interface{}
}

type CreateContentResponse struct {
	AppID   string
	Content *Content
}

type DeleteContentRequest struct {
	AppID     string
	ContentID int64
}

type DeleteResponse struct {
	AppID     string
	ContentID int64
}
type Service interface {
	GetContent(ctx context.Context, req *GetContentRequest) (*GetContentResponse, error)
	ListContents(ctx context.Context, req *ListContentsRequest) (*ListContentsResponse, error)
	CreateContent(ctx context.Context, req *CreateContentRequest) (*CreateContentResponse, error)
	DeleteContent(ctx context.Context, req *DeleteContentRequest) (*DeleteResponse, error)
}
