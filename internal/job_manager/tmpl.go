package job_manager

import (
	log "github.com/sirupsen/logrus"
)

type App struct {
	config *Config
}

func NewApp(config *Config) (*App, error) {
	return &App{config}, nil
}

func (a *App) Run() error {
	log.Info("Running application")

	repo, err := OpenRepo(a.config.PostgresAddress)
	if err != nil {
		log.WithError(err).Errorf("Failed to open repo")
		return err
	}

	server, err := CreateServer(a.config, repo)
	if err != nil {
		log.WithError(err).Errorf("Failed to create server")
		return err
	}

	return server.Run()
}
