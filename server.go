package main

import (
	"github.com/sundogrd/content-grpc/repositories/content/repo"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"

	"github.com/sirupsen/logrus"

	contentGen "github.com/sundogrd/content-grpc/grpc_gen/content"
	"github.com/sundogrd/content-grpc/servers/content"
	configUtils "github.com/sundogrd/gopkg/config"
	dbUtils "github.com/sundogrd/gopkg/db"
	grpcUtils "github.com/sundogrd/gopkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := configUtils.Load("./config", "service.config")
	if err != nil {
		logrus.Errorf("[content-grpc] ReadConfigFromFile err: %s", err.Error())
		panic(err)
	}

	instanceAddr := config.Get("grpcService.host").(string) + ":" + config.Get("grpcService.port").(string)
	listen, err := net.Listen("tcp", instanceAddr)
	if err != nil {
		logrus.Errorf("[content-grpc] net.Listen err: %s", err.Error())
		panic(err)
	}

	gormDB, err := dbUtils.Connect(dbUtils.ConnectOptions{
		User:           config.Get("db.options.user").(string),
		Password:       config.Get("db.options.password").(string),
		Host:           config.Get("db.options.host").(string),
		Port:           int16(config.Get("db.options.port").(float64)),
		DBName:         config.Get("db.options.dbname").(string),
		ConnectTimeout: config.Get("db.options.connectTimeout").(string),
	})

	if err != nil {
		logrus.Errorf("[content-grpc] db.Connect err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[content-grpc] db.Connect finished")

	contentRepo, err := repo.NewContentRepo(gormDB, 2 * time.Second)
	if err != nil {
		logrus.Errorf("[content-grpc] NewContentRepo err=%s", err.Error())
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 1 * time.Minute,           // <--- This fixes it!
		}),
	)

	endpoints := config.Get("etcd.host").(string) + ":" + config.Get("etcd.port").(string)
	logrus.Printf("etcd config is %s", endpoints)

	resolver, err := grpcUtils.NewGrpcResolover(endpoints)
	if err != nil {
		logrus.Errorf("[content-grpc] NewGrpcResolover err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[content-grpc] NewGrpcResolover finished")

	serverName := config.Get("etcd.serverName").(string)
	err = grpcUtils.ResgiterServer(*resolver, serverName, instanceAddr, 5*time.Second, 5)
	if err != nil {
		logrus.Errorf("[content-grpc] RegisterServer err: %s", err.Error())
		panic(err)
	}
	logrus.Printf("[content-grpc] ResgiterServer finished, service: %s, %s", serverName, instanceAddr)

	contentGen.RegisterContentServiceServer(grpcServer, &content.ContentServiceServer{
		GormDB: gormDB,
		ContentRepo: contentRepo,
	})
	reflection.Register(grpcServer)
	grpcServer.Serve(listen)
}
