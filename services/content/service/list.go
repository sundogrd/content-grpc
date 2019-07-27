package service

import (
	"context"
	"github.com/sirupsen/logrus"
	contentRepo "github.com/sundogrd/content-grpc/repositories/content"
	service "github.com/sundogrd/content-grpc/services/content"
)

func (s *contentService) ListContent(ctx context.Context, req *service.ListContentsRequest) (*service.ListContentsResponse, error) {
	page := int32(1)
	pageSize := int32(10)
	if req.Page != 0 {
		page = req.Page
	}
	if req.PageSize != 0 {
		pageSize = req.PageSize
	}

	repoListResp, err := s.contentRepo.List(ctx, &contentRepo.ListRequest{
		Page:     &page,
		PageSize: &pageSize,
	})
	if err != nil {
		logrus.Errorf("[service/comment] List: list errored: %+v", err)
		return nil, err
	}

	var ret []*service.Content
	for _, item := range repoListResp.List {
		ret = append(ret, &service.Content{
			ContentID:   item.ContentID,
			AppID:       item.AppID,
			Title:       item.Title,
			Description: item.Description,
			AuthorID:    item.AuthorID,
			Category:    item.Category,
			Type:        item.Type,
			Body:        item.Body,
			BodyType:    item.BodyType,
			State:       item.State,
			Extra:       item.Extra,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	res := &service.ListContentsResponse{
		AppID:    req.AppID,
		Contents: ret,
		Page:     repoListResp.Page,
		PageSize: repoListResp.PageSize,
		Total:    repoListResp.Total,
	}

	return res, nil
}
