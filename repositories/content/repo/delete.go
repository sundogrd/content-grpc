package repo

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-grpc/repositories"

	repo "github.com/sundogrd/content-grpc/repositories/content"
)

func (r contentRepo) Delete(ctx context.Context, req *repo.DeleteRequest) (*repo.DeleteResponse, error) {
	db := r.gormDB

	var content repositories.Content
	db.Where("content_id = ?", req.ContentID).First(&content)

	content.State = repositories.STATE_DELETED

	if dbc := db.Save(&content); dbc.Error != nil {
		logrus.Errorf("[repository/content] Delete: db delete error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.DeleteResponse{
		ContentID: req.ContentID,
	}

	return res, nil
}
