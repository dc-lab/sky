package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"

	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/data_manager/master/config"
	"github.com/dc-lab/sky/internal/data_manager/master/repo"
	"github.com/dc-lab/sky/internal/data_manager/master/server"
	"github.com/dc-lab/sky/internal/data_manager/master/service"
)

func initLogging(conf *config.Config) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	log.SetReportCaller(true)

	if conf.LogFile != "" {
		log.WithField("file", conf.LogFile).Info("Initializing log file")
		logFile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.WithError(err).Fatalln("Failed to open log file")
		}
		mw := io.MultiWriter(os.Stderr, logFile)
		log.SetOutput(mw)
	}
}

func runGrpcGateway(wg *sync.WaitGroup, srv *service.FilesService) {
	defer wg.Done()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := gw_runtime.NewServeMux(
		gw_runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "User-Id":
				return key, true
			default:
				return gw_runtime.DefaultHeaderMatcher(key)
			}
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterDataManagerHandlerFromEndpoint(ctx, mux, srv.Config.GrpcBindAddress, opts)
	if err != nil {
		log.WithError(err).Fatalln("Failed to start grpc-gateway")
	}

	log.Println("Starting grpc-gateway server at", srv.Config.HttpBindAddress)
	err = http.ListenAndServe(srv.Config.HttpBindAddress, mux)
	log.WithError(err).Fatalln("Grpc-gateway server failed")
}

func runGrpcServer(wg *sync.WaitGroup, srv *service.FilesService) {
	defer wg.Done()

	server := &server.Server{Files: srv}

	s := grpc.NewServer()
	pb.RegisterDataManagerServer(s, server)
	pb.RegisterMasterServer(s, server)
	reflection.Register(s)

	log.Println("Starting grpc server at", srv.Config.GrpcBindAddress)
	lis, err := net.Listen("tcp", srv.Config.GrpcBindAddress)
	if err != nil {
		log.WithError(err).Fatalln("Failed to listen socket for grpc server")
	}
	err = s.Serve(lis)
	log.WithError(err).Fatalln("Grpc server failed")
}

// @title Sky Data Manager
// @version 1.0
// @description This is data manager for Sky platform.

// @contact.name @BigRedEye
// @contact.url https://t.me/BigRedEye
// @contact.email mail@sskvor.dev

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host sky.sskvor.dev
// @BasePath /v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Info("Failed to load config from .env file")
	}

	conf, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatalln("Failed to load config")
	}

	initLogging(conf)

	repo, err := repo.OpenFilesRepo("postgres", conf.PostgresAddress)
	if err != nil {
		log.WithError(err).Fatalln("Could not connect to database")
	}

	srv := service.FilesService{
		Repo:   repo,
		Config: conf,
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go runGrpcServer(&wg, &srv)
	go runGrpcGateway(&wg, &srv)

	wg.Wait()
}
