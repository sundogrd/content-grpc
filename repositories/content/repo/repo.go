package repo

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sundogrd/content-grpc/repositories/content"
	"github.com/zheng-ji/goSnowFlake"
)

type contentRepo struct {
	gormDB         *gorm.DB
	contextTimeout time.Duration
	idBuilder      *goSnowFlake.IdWorker
}

// NewUserService will create new an articleUsecase object representation of article.Usecase interface
func NewContentRepo(gormDB *gorm.DB, timeout time.Duration) (content.Repo, error) {
	idBuilder, err := goSnowFlake.NewIdWorker(3)
	if err != nil {
		return nil, err
	}

	hasTable := gormDB.HasTable(&content.Content{})
	if hasTable == false {
		gormDB.CreateTable(&content.Content{})
	} else {
		gormDB.AutoMigrate(&content.Content{})
	}

	repo := contentRepo{
		gormDB:         gormDB,
		contextTimeout: timeout,
		idBuilder:      idBuilder,
	}
	return repo, nil
}
