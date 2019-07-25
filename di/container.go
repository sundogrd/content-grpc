package di

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/sundogrd/gopkg/config"
	"github.com/sundogrd/gopkg/db"
)

type Container struct {
	Config *viper.Viper
	SundogDB *gorm.DB
}

func InitContainer() (*Container, error) {
	conf, err := config.Load("./config", "service.config")
	if err != nil {
		return nil, err
	}

	sundogDB, err := db.Connect(db.ConnectOptions{
		User: conf.GetString("db.options.user"),
		Password: conf.GetString("db.options.passport"),
		Host: conf.GetString("db.options.host"),
		Port: conf.GetString("db.options.port"),
		DBName: conf.GetString("db.options.dbname"),
		ConnectTimeout: conf.GetString("db.options.connectTimeout"),
	})


	return &Container{
		Config: conf,
		SundogDB: sundogDB,
	}, nil
}
