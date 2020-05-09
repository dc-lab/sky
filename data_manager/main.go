package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"

	"data_manager/config"
	"data_manager/handlers"
	"data_manager/repo"
	"data_manager/router"
)

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
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	log.SetReportCaller(true)

	log.Info("Loading config")
	conf, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatalln("Failed to laod config")
	}

	if conf.LogFile != "" {
		log.WithField("file", conf.LogFile).Info("Initializing log file")
		logFile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			log.WithError(err).Fatalln("Failed to open log file")
		}
		mw := io.MultiWriter(os.Stderr, logFile)
		log.SetOutput(mw)
	}

	err = os.Mkdir(conf.StorageDir, 0o755)
	if err != nil && !os.IsExist(err) {
		log.WithError(err).WithField("dir", conf.StorageDir).Fatalln("Failed to create storage directory")
	}

	repo, err := repo.OpenFilesRepo("postgres", conf.PostgresAddress)
	if err != nil {
		log.WithError(err).Fatalln("Could not connect to database")
	}
	if err := repo.Migrate(); err != nil {
		log.WithError(err).Fatalln("Failed to run migrations")
	}

	srv := handlers.FilesService{
		Repo:   repo,
		Config: conf,
	}

	handler := routers.MakeRouter(srv)

	log.Println("Starting server at", conf.BindAddress)
	if err := http.ListenAndServe(conf.BindAddress, handler); err != nil {
		log.WithError(err).Fatalln("Main server process failed")
	}
}
