package main

import (
	"github.com/dc-lab/sky/agent/src/network"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err)
	}
	network.RunClient()
}
