package main

import (
	"fmt"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/dc-lab/sky/internal/gateway"
	"github.com/dc-lab/sky/pkg/logging"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Failed to load .env file")
	}

	conf, err := gateway.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	logging.InitLogging(logging.Config{
		LogFile:  conf.LogFile,
		LogLevel: "info",
	})

	app, err := gateway.NewApp(conf)
	if err != nil {
		return fmt.Errorf("failed to initialize app: %w", err)
	}

	return app.Run()
}

func main() {
	err := run()
	if err != nil {
		log.WithError(err).Fatalln("gateway process failed")
	}
}
