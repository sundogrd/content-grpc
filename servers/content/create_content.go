package content

import (
	"context"
	"errors"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
)

func (server ContentServiceServer) CreateContent(ctx context.Context, req *content.CreateContentRequest) (*content.CreateContentResponse, error) {
	return nil, errors.New("not implemented")
}
