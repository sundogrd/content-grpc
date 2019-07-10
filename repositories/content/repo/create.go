package repo

import (
	"context"
	"github.com/sirupsen/logrus"
	repo "github.com/sundogrd/content-grpc/repositories/content"
)

// 只负责创建
func (r contentRepo) Create(ctx context.Context, req *repo.CreateRequest) (*repo.CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	contentID, _ := r.idBuilder.NextId()

	// 支持通过Status创建，默认为Published
	contentState := repo.STATE_PUBLISHED

	contentDescription := ""
	if req.Description != nil {
		contentDescription = *req.Description
	}

	contentCategory := ""
	if req.Category != nil {
		contentCategory = *req.Category
	}

	contentType := repo.TYPE_RICHTEXT
	if req.Type != nil {
		contentType = *req.Type
	}

	contentBodyType := repo.BODY_TYPE_HTML
	if req.BodyType != nil {
		contentBodyType = *req.BodyType
	}

	content := repo.Content{
		AppID: req.AppId,
		ContentID:   contentID,
		Title:       req.Title,
		Description: contentDescription,
		AuthorID:    req.AuthorID,
		Category:    contentCategory,
		State:       contentState,
		Type:        contentType, // 先写死只有图文
		Body:        req.Body,
		BodyType:    contentBodyType,
		Extra:       "{}",
	}
	if dbc := r.gormDB.Create(&content); dbc.Error != nil {
		logrus.Errorf("[content-grpc/providers/repos/content] Create: content error: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	return &repo.CreateResponse{
		AppId: req.AppId,
		Content: &content,
	}, nil
}
