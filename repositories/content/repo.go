package content

import (
	"context"
)

type GetRequest struct {
	ContentId int64
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
	AppId   string
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
	AppId   string
	Content *Content
}

type DeleteRequest struct {
	ContentId int64
}

type DeleteResponse struct {
	ContentId int64
}

type UpdateRequest struct {
	ContentId int64
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

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}
