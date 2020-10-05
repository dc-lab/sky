package main

import (
	"fmt"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	jm "github.com/dc-lab/sky/internal/job_manager"
	"github.com/dc-lab/sky/pkg/logging"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Failed to load .env file")
	}

	conf, err := jm.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = logging.InitLogging(logging.Config{
		LogFile:  conf.LogFile,
		LogLevel: "info",
	})
	if err != nil {
		return fmt.Errorf("failed to initialize logging: %w", err)
	}

	app, err := jm.NewApp(conf)
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	return app.Run()
}

func main() {
	err := run()
	if err != nil {
		log.WithError(err).Fatalln("Main tmpl process failed")
	}
}
