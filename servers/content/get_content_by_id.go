package content

import (
	"context"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
)

func (server ContentServiceServer) GetContentById(ctx context.Context, req *content.GetContentByIdRequest) (*content.GetContentByIdResponse, error) {

	return &content.GetContentByIdResponse{
	}, nil
}
