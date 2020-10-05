package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/dc-lab/sky/internal/data_manager/node/config"
	"github.com/dc-lab/sky/internal/data_manager/node/router"
	"github.com/dc-lab/sky/internal/data_manager/node/service"
	"github.com/dc-lab/sky/internal/data_manager/node/storage"
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

func runHttpServer(wg *sync.WaitGroup, srv *service.BlobsService, conf *config.Config) {
	defer wg.Done()

	handler := router.MakeRouter(srv)

	log.Println("Starting http server at ", conf.HttpBindAddress)
	err := http.ListenAndServe(conf.HttpBindAddress, handler)
	log.WithError(err).Fatalln("Http server failed")
}

func runNodeLoop(wg *sync.WaitGroup, srv *service.BlobsService, conf *config.Config) {
	defer wg.Done()

	log.Println("Starting node loop with master at ", conf.MasterAddress)

	srv.RunLoop()
}

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

	storage, err := storage.MakeLocalStorage(conf)
	if err != nil {
		log.WithError(err).Fatalln("Failed to initialize local storage")
	}

	service, err := service.MakeBlobsService(conf, storage)
	if err != nil {
		log.WithError(err).Fatalln("Failed to initialize local storage")
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go runHttpServer(&wg, service, conf)

	wg.Add(1)
	go runNodeLoop(&wg, service, conf)

	wg.Wait()
}
