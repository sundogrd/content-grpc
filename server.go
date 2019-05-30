package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"time"

	configUtils "github.com/sundogrd/gopkg/config"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	contentGen "github.com/sundogrd/content-grpc/grpc_gen/content"
	"github.com/sundogrd/content-grpc/servers/content"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configUtils.ReadConfigFromFile("./config", nil)
	if err != nil {
		logrus.Errorf("[search-grpc] ReadConfigFromFile err: %s", err.Error())
		panic(err)
	}

	instanceAddr := config.Get("grpcService.host").(string) + ":" + config.Get("grpcService.port").(string)
	listen, err := net.Listen("tcp", instanceAddr)
	if err != nil {
		logrus.Errorf("[search-grpc] net.Listen err: %s", err.Error())
		panic(err)
	}

	gormDB, err := db.Connect(db.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           config.Get("db.options.port").(string),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})
	if err != nil {
		logrus.Errorf("[search-grpc] db.Connect err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] db.Connect finished")

	grpcServer := grpc.NewServer()
	resolver, err := grpcUtils.NewGrpcResolover()
	if err != nil {
		logrus.Errorf("[search-grpc] NewGrpcResolover err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] NewGrpcResolover finished")

	err = grpcUtils.ResgiterServer(*resolver, "sundog.search", instanceAddr, 5*time.Second, 5)
	if err != nil {
		logrus.Errorf("[search-grpc] RegisterServer err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[search-grpc] ResgiterServer finished, service: %s, %s", "sundog.search", instanceAddr)

	contentGen.RegisterContentServiceServer(grpcServer, &content.ContentServiceServer{
		GormDB:              gormDB,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
