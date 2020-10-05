package main

import (
	"fmt"

	agent "github.com/dc-lab/sky/internal/agent/src"
	logging "github.com/dc-lab/sky/pkg/logging"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Failed to load .env file")
	}

	config, err := agent.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	err = logging.InitLogging(logging.Config{
		LogFile:  config.LogFile,
		LogLevel: config.LogLevel,
	})
	if err != nil {
		return err
	}

	app, err := agent.NewApp(config)
	if err != nil {
		return fmt.Errorf("failed to create app: %w", err)
	}

	return app.Run()
}

func main() {
	err := run()
	if err != nil {
		log.WithError(err).Fatalf("Agent failed")
	}
}
