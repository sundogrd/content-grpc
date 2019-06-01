package content

import (
	"context"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
)

func (server ContentServiceServer) ListContents(ctx context.Context, req *content.ListContentsRequest) (*content.ListContentsResponse, error) {

	return &content.ListContentsResponse{
	}, nil
}
