package main

import (
	"context"
	"main/database"
	"main/models"
	"main/server"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	var serverLogFilePath = "./logs/server.log"

	if file, err := os.OpenFile(serverLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766); err != nil {
		logrus.Errorf("can't open file '%s' (cause: %s) - logging into STDOUT", serverLogFilePath, err.Error())
	} else {
		logrus.SetOutput(file)
		logrus.Info("logging intited")
	}
}

const (
	port = ":50051"
)

func main() {

	// db, err := database.InitConnection("local.db")
	db, err := database.InitConnection("file::memory:")
	if err != nil {
		logrus.Fatal("Failed to connect database 'local.db'")
	}

	ctx := context.Background()

	s, err := server.New(ctx, db, 30*time.Second)
	if err != nil {
		logrus.Fatalf("Failed initialize server '%s'", err.Error())
	}

	defer func() {
		if _, err := s.Stop(ctx, nil); err != nil {
			logrus.Errorf("error while stopping server (cause '%s')", err.Error())
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	models.RegisterRssServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
