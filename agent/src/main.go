package main

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/network"
	"github.com/dc-lab/sky/agent/src/parser"
	"log"
	"os"
	"path"
)

func initLogs() *os.File {
	logPath := path.Join(parser.AgentConfig.LogsDirectory, "agent.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	common.DieWithError(err)
	log.SetOutput(file)
	return file
}

func main() {
	logFile := initLogs()
	defer logFile.Close()
	network.RunClient()
}
