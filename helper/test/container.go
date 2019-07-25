package test

import (
	"github.com/sundogrd/content-grpc/di"
	"github.com/sundogrd/gopkg/db"
)

func InitTestContainer() (*di.Container, error) {
	sundogDB, err := db.Connect(db.ConnectOptions{
		User: "sundog",
		Password: "123456",
		Host: "localhost",
		Port: "3306",
		DBName: "sundog",
		ConnectTimeout: "10s",
	})

	if err != nil {
		return nil, err
	}

	return &di.Container{
		Config: nil,
		SundogDB: sundogDB,
	}, nil
}
