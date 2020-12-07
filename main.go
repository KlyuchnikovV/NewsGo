package main

import (
	"context"
	"fmt"
	"main/database"
	"main/models"
	"main/server"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	var serverLogFilePath = "./server.log"

	if file, err := os.OpenFile(serverLogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0766); err != nil {
		logrus.Errorf("can't open file '%s' (cause: %s) - logging into STDOUT", serverLogFilePath, err.Error())
	} else {
		logrus.SetOutput(file)
		logrus.Info("logging intited")
	}

	if len(os.Args) > 1 {
		if err := parseArguments(os.Args[1:]); err != nil {
			logrus.Fatalf("Can't parse arguments (cause: %s)", err.Error())
		}
	}
}

var (
	dbPath     string = "file::memory:"
	port       int64  = 50051
	runOnStart bool   = false
)

func main() {
	logrus.Infof("Db file is %s", dbPath)
	db, err := database.InitConnection(dbPath)
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

	if runOnStart {
		logrus.Info("Starting collectors")
		if _, err := s.Start(ctx, nil); err != nil {
			logrus.Fatalf("Failed to start: %v", err)
		}
	}

	logrus.Infof("Listening port :%d", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	models.RegisterRssServer(grpcServer, s)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}

// Format "--<key>=<value>"
func parseArguments(args []string) error {
	r := regexp.MustCompile(`^--([a-z-]+)=([a-z0-9\.]+)$`)
	for _, arg := range args {
		var err error
		values := r.FindAllStringSubmatch(arg, 1)
		if len(values[0]) != 1 && len(values[0]) != 3 {
			return fmt.Errorf("wrong argument %s", arg)
		}
		switch values[0][1] {
		case "run-on-start":
			runOnStart, err = strconv.ParseBool(values[0][2])
		case "ram-db":
			dbPath = values[0][2]
		case "port":
			port, err = strconv.ParseInt(values[0][2], 10, 64)
			if err != nil {
				break
			}
			if port < 0 || port > 65536 {
				err = fmt.Errorf("wrong port defined %d", port)
			}
		default:
			err = fmt.Errorf("unknown argument %s", arg)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
