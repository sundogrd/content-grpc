package content

import (
	"context"
)

type GetRequest struct {
	ContentID int64
}
type GetResponse struct {
	Content *Content
}

type ListRequest struct {
	Page     int32
	PageSize int32
}
type ListResponse struct {
	List     []*Content
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
	Type     *ContentType
	Body     string
	BodyType *ContentBodyType
	Extra    map[string]interface{}

}
type CreateResponse struct {
	AppID   string
	Content *Content
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
	State *ContentState
	Type *ContentType
	Body *string
	BodyType *ContentBodyType
	Extra *string
}

type UpdateResponse struct {
	Content *Content
}
I
type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}
