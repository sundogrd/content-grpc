package content_test

import (
	"github.com/sundogrd/content-grpc/repositories/content/repo"
	"github.com/sundogrd/content-grpc/servers/content"
	"github.com/sundogrd/gopkg/db"
	"time"
)

func newTestContentServer() (*content.ContentServiceServer, error) {
	gormDB, err := db.Connect(db.ConnectOptions{
		User:           "sundog",
		Password:       "sundogPwd",
		Host:           "127.0.0.1",
		Port:           3306,
		DBName:         "sundog",
		ConnectTimeout: "10s",
	})
	if err != nil {
		return nil, err
	}
	contentRepo, err := repo.NewContentRepo(gormDB, 2 * time.Second)
	return &content.ContentServiceServer{
		GormDB: gormDB,
		ContentRepo: contentRepo,
	}, nil
}
