package content

import (
	"context"
	"github.com/sundogrd/content-grpc/repositories"
)

type GetRequest struct {
	ContentID int64
}
type GetResponse struct {
	Content *repositories.Content
}

type ListRequest struct {
	Page     *int32
	PageSize *int32
	Query    string
}
type ListResponse struct {
	List     []*repositories.Content
	Page     int32
	PageSize int32
	Total    int64
}

type CreateRequest struct {
	AppID   string
	Title   string
	Description *string
	AuthorID int64
	Category *string
	Type     *repositories.ContentType
	Body     string
	BodyType *repositories.ContentBodyType
	Extra    map[string]interface{}

}
type CreateResponse struct {
	AppID   string
	Content *repositories.Content
}

type DeleteRequest struct {
	ContentID int64
}

type DeleteResponse struct {
	ContentID int64
}

type UpdateRequest struct {
	ContentID int64
	Title *string
	Description *string
	Category *string
	State *repositories.ContentState
	Type *repositories.ContentType
	Body *string
	BodyType *repositories.ContentBodyType
	Extra *string
}

type UpdateResponse struct {
	Content *repositories.Content
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}
