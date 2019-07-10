package service
//
//import (
//	"github.com/sundogrd/content-grpc/services/content"
//	"github.com/zheng-ji/goSnowFlake"
//	"time"
//
//	contentRepo "github.com/sundogrd/content-grpc/providers/repos/content"
//)
//
//type contentService struct {
//	contentRepo    contentRepo.Repo
//	contextTimeout time.Duration
//	idBuilder *goSnowFlake.IdWorker
//	content.Service
//}
//
//// NewCommentService will create new an articleUsecase object representation of article.Usecase interface
//func NewContentService(contentRepo contentRepo.Repo, timeout time.Duration) (content.Service, error) {
//	idBuilder, err := goSnowFlake.NewIdWorker(443474713)
//	if err != nil {
//		return nil, err
//	}
//	return &contentService{
//		contentRepo:    contentRepo,
//		contextTimeout: timeout,
//		idBuilder: idBuilder,
//	}, nil
//}
