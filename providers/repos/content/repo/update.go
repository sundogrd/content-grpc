package repo

import (
	"context"
	"errors"
	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func (r contentRepo) Update(ctx context.Context, req *repo.UpdateRequest) (*repo.UpdateResponse, error) {

	//db := r.gormDB

	//dbc := db.Model(&repo.Content{}).Where("id = ?", req.ContentId).Update(req.Map)
	//
	//if dbc.Error != nil {
	//	fmt.Printf("[providers/comment] Update: db update error: %+v", dbc.Error)
	//	return nil, dbc.Error
	//}
	//
	//response, err := r.Get(ctx, &repo.GetRequest{
	//	ContentId: req.ContentId,
	//})
	//
	//if err != nil {
	//	fmt.Printf("[providers/comment] Update: db get comment after update error: %+v", err)
	//	return nil, err
	//}
	//
	//res := &repo.UpdateResponse{
	//	Comment: response.Comment,
	//}

	//return res, nil

	return nil, errors.New("not implemented")

}
