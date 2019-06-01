package repo

import (
	"context"
	"github.com/sirupsen/logrus"

	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func (r contentRepo) Delete(ctx context.Context, req *repo.DeleteRequest) (*repo.DeleteResponse, error) {
	db := r.gormDB

	var content repo.Content
	db.Where("content_id = ?", req.ContentId).First(&content)

	content.State = repo.STATE_DELETED

	if dbc := db.Save(&content); dbc.Error != nil {
		logrus.Errorf("[providers/comment] Delete: db delete error: %+v", dbc.Error)
		return nil, dbc.Error
	}

	res := &repo.DeleteResponse{
		ContentId: req.ContentId,
	}

	return res, nil
}
