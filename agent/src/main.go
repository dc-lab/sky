package main

import (
	"github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/network"
	"github.com/dc-lab/sky/agent/src/parser"
	"log"
	"os"
	"path"
)

func initLogs() {
	logPath := path.Join(parser.AgentConfig.AgentDirectory, "sky-agent.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	common.DieWithError(err)
	defer file.Close()
	log.SetOutput(file)
}

func main() {
	initLogs()
	network.RunClient()
}
