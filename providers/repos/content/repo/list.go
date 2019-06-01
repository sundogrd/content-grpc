package repo

import (
	"context"
	"errors"

	repo "github.com/sundogrd/content-grpc/providers/repos/content"
)

func (s contentRepo) List(ctx context.Context, req *repo.ListRequest) (*repo.ListResponse, error) {
	//var page int16 = 1
	//var pageSize int16 = 10
	//if req.Page != nil {
	//	page = *req.Page
	//}
	//if req.PageSize != nil {
	//	pageSize = *req.PageSize
	//}
	//
	//contents := make([]SDContent, 0)
	//count := int64(0)
	//
	//db := cs.db
	//if req.ContentIDs != nil && len(*req.ContentIDs) != 0 {
	//	db = db.Where("content_id in (?)", *req.ContentIDs)
	//}
	//if req.Title != nil {
	//	db = db.Where("title LIKE ?", "%"+*req.Title+"%")
	//}
	//if req.AuthorID != nil {
	//	db = db.Where("author_id = ", *req.AuthorID)
	//}
	//if req.Type != nil {
	//	db = db.Where("type = ?", *req.Type)
	//}
	//if req.Status != nil {
	//	db = db.Where("status = ?", *req.Status)
	//}
	//db.Limit(pageSize).Offset((page - 1) * (pageSize))
	//if err := db.Find(&contents).Offset(0).Limit(-1).Count(&count).Error; err != nil {
	//	return nil, err
	//} else {
	//	BaseInfos := make([]BaseInfo, 0)
	//	for _, v := range contents {
	//		BaseInfos = append(BaseInfos, packBaseInfo(v))
	//	}
	//	res := &FindResponse{
	//		List:  BaseInfos,
	//		Total: count,
	//	}
	//	return res, nil
	//}

	//return res, nil

	return nil, errors.New("not implemented")
}
