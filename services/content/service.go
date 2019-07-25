package content

import (
	"context"
	"github.com/sundogrd/content-grpc/repositories/content"

	"github.com/sundogrd/comment-grpc/models"
)

type ECommentState int32

const (
	UNKNOWN  ECommentState = 0
	SHOW     ECommentState = 1
	WITHDRAW ECommentState = 2
)

type GetRequest struct {
	CommentId int64
}
type GetResponse struct {
	Comment *models.Comment
}

type ListContentRequest struct {
	AppID       string
	TargetId    int64
	Page        int32
	PageSize    int32
	CreatorId   int64
	ParentId    int64
	State       int32
	StartTime   uint32
	EndTime     uint32
	ReCommentId int64
}

type ListContentResponse struct {
	AppID    string
	Contents []*Content
	Total    int64
	Page     int32
	PageSize int32
}

type CreateRequest struct {
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

type CreateResponse struct {
	AppID   string
	ContentID int64
}

type DeleteRequest struct {
	AppID     string
	ContentID int64
}

type DeleteResponse struct {
	AppID     string
	ContentID int64
}
type Service interface {
	GetComment(ctx context.Context, req *GetRequest) (*GetResponse, error)
	ListComments(ctx context.Context, req *ListContentRequest) (*ListContentResponse, error)
	CreateComment(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	DeleteComment(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
}
