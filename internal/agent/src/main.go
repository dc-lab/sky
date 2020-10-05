package agent

import (
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/dc-lab/sky/internal/agent/src/network"
	"github.com/dc-lab/sky/internal/agent/src/parser"
)

func initLogs(logFilePath string) *os.File {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	common.DieWithError(err)
	log.SetOutput(file)
	return file
}

func LoadConfig() (*parser.Config, error) {
	opts := parser.ParseArguments()

	// FIXME: Agent stores config in the global variable, refactor this
	config, err := parser.InitializeAgentConfigFromOptions(&opts)
	if err != nil {
		return nil, err
	}

	return config, nil
}

type App struct {
	Config *parser.Config
	Client *network.Client
}

func NewApp(config *parser.Config) (*App, error) {
	client, err := network.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &App{config, client}, nil
}

func (a *App) Run() error {
	// FIXME: Agent panics on any error
	return a.Client.Run()
}
