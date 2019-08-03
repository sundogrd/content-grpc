package repo

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-grpc/repositories"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func (r contentRepo) Get(ctx context.Context, req *repo.GetRequest) (*repo.GetResponse, error) {

	var content repositories.Content
	db := r.gormDB

	dbc := db.Where(repositories.Content{
		ContentID: req.ContentID,
	}).First(&content)

	if dbc.Error != nil {
		logrus.Errorf("[repositories/content] Delete: db get error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.GetResponse{
		Content: &content,
	}

	if res.Content.ID == 0 {
		return nil, nil
	}

	return res, nil
}
