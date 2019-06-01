package content

import (
	"context"
	"github.com/sundogrd/content-grpc/grpc_gen/content"
)

func (server ContentServiceServer) DeleteContent(ctx context.Context, req *content.DeleteContentRequest) (*content.DeleteContentResponse, error) {

	return &content.DeleteContentResponse{
	}, nil
}
