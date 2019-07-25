package service

import (
	"context"
	"github.com/sirupsen/logrus"
	contentRepo "github.com/sundogrd/content-grpc/repositories/content"
	service "github.com/sundogrd/content-grpc/services/content"
)

func (s *contentService) Create(ctx context.Context, req *service.CreateRequest) (*service.CreateResponse, error) {

	repoCreateResponse, err := s.contentRepo.Create(ctx, &contentRepo.CreateRequest{
		AppID: req.AppID,
		Title:   req.Title,
		Description: req.Description,
		AuthorID: req.AuthorID,
		Category: req.Category,
		Type:    req.Type,
		Body:     req.Body,
		BodyType: req.BodyType,
		Extra: req.Extra,
	})
	if err != nil {
		logrus.Errorf("[service/comment] List: list errored: %+v", err)
		return nil, err
	}

	res := &service.CreateResponse{
		AppID:    req.AppID,
		ContentID: repoCreateResponse.Content.ContentID,
	}

	return res, nil
}
