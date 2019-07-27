package service

import (
	"context"
	"github.com/sirupsen/logrus"
	contentRepo "github.com/sundogrd/content-grpc/repositories/content"
	service "github.com/sundogrd/content-grpc/services/content"
)

func (s *contentService) CreateContent(ctx context.Context, req *service.CreateContentRequest) (*service.CreateContentResponse, error) {

	createResp, err := s.contentRepo.Create(ctx, &contentRepo.CreateRequest{
		AppID:       req.AppID,
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    req.AuthorID,
		Category:    req.Category,
		Type:        req.Type,
		Body:        req.Body,
		BodyType:    req.BodyType,
		Extra:       req.Extra,
	})
	if err != nil {
		logrus.Errorf("[service/comment] List: list errored: %+v", err)
		return nil, err
	}

	res := &service.CreateContentResponse{
		AppID:    req.AppID,
		Content: &service.Content{
			ContentID: createResp.Content.ContentID,
			AppID: createResp.Content.AppID,
			State: createResp.Content.State,
			AuthorID: createResp.Content.AuthorID,
			Body: createResp.Content.Body,
			BodyType: createResp.Content.BodyType,
			Extra: createResp.Content.Extra,
			Type: createResp.Content.Type,
			Category: createResp.Content.Category,
			Description: createResp.Content.Description,
			UpdatedAt: createResp.Content.UpdatedAt,
			CreatedAt: createResp.Content.CreatedAt,
			Title: createResp.Content.Title,
		},
	}

	return res, nil
}
