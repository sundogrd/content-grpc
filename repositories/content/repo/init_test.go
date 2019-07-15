package repo_test

import (
	"time"

	repo "github.com/sundogrd/content-grpc/repositories/content"
	contentRepo "github.com/sundogrd/content-grpc/repositories/content/repo"
	"github.com/sundogrd/gopkg/db"
)

func initTestRepo() (repo.Repo, error) {

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           "sundog",
		Password:       "sundogPwd",
		Host:           "127.0.0.1",
		Port:           "3306",
		DBName:         "sundog",
		ConnectTimeout: "10s",
	})
	if err != nil {
		return nil, err
	}
	content, error := contentRepo.NewContentRepo(gormDB, 2*time.Second)
	if error != nil {
		return nil, error
	}
	return content, nil
}
