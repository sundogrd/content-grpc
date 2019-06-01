package service

import (
	"context"
	service "github.com/sundogrd/content-grpc/services/content"
	contentRepo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func (s *contentService) CreateContent(ctx context.Context, req *service.CreateRequest) (*service.CreateResponse, error) {
	res, err := s.contentRepo.Create(context.Background(), &contentRepo.CreateRequest{

	})
	if err != nil {
		return nil, err
	}
	return &service.CreateResponse{

	}, nil
}
