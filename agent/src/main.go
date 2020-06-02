package main

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/network"
	"github.com/dc-lab/sky/agent/src/parser"
	"log"
	"os"
)

func initLogs(logFilePath string) *os.File {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	common.DieWithError(err)
	log.SetOutput(file)
	return file
}

func main() {
	logFile := initLogs(parser.AgentConfig.AgentLogFile)
	defer logFile.Close()
	network.RunClient()
}
