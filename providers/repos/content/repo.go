package content

import (
	"context"
)

type GetRequest struct {
	CommentId int64
}
type GetResponse struct {
	Comment *Comment
}

type ListRequest struct {
	Query    string
	Page     int32
	PageSize int32
	// Receiver interface{}
	Values []interface{}
}
type ListResponse struct {
	List     []*Comment
	Page     int32
	PageSize int32
	Total    int64
}

type CreateRequest struct {
	AppId   string
	Comment CommentParams
}
type CreateResponse struct {
	AppId   string
	Comment *Comment
}

type DeleteRequest struct {
	CommentId int64
}

type DeleteResponse struct {
	CommentId int64
}

type UpdateRequest struct {
	CommentId int64
	Map       map[string]interface{}
}

type UpdateResponse struct {
	Comment *Comment
}

type Repo interface {
	Get(ctx context.Context, req *GetRequest) (*GetResponse, error)
	List(ctx context.Context, req *ListRequest) (*ListResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error)
	Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error)
}
