package parser

import (
	"flag"
	"fmt"
	common "github.com/dc-lab/sky/agent/src/common"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
)

type Config struct {
	ResourceManagerAddress string
	AgentDirectory         string
	AgentLogFile           string
	HealthFile             string
	Token                  string
}

func GetToken(path string) string {
	bytes, err := ioutil.ReadFile(path)
	common.DieWithError(err)
	return string(bytes)
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigType("json")
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func ParseArguments() Config {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "Path to agent configuration file")
	flag.Parse()

	v, err := readConfig(configPath, map[string]interface{}{
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/var/tmp/agent",
		"LogsDirectory":          "/var/tmp/agent-logs",
		"RunDirectory":           "/var/run/agent",
		"TokenPath":              "/var/tmp/token",
	})
	common.DieWithError(err)

	token := GetToken(v.GetString("TokenPath"))
	logsDirectory := v.GetString("LogsDirectory")
	runDirectory := v.GetString("RunDirectory")
	config := Config{
		ResourceManagerAddress: v.GetString("ResourceManagerAddress"),
		AgentDirectory:         v.GetString("AgentDirectory"),
		AgentLogFile:           path.Join(logsDirectory, "agent.log"),
		HealthFile:             path.Join(runDirectory, "health.info"),
		Token:                  token,
	}

	common.DieWithError(common.CreateDirectory(config.AgentDirectory, true))
	common.DieWithError(common.CreateDirectory(logsDirectory, false))
	common.DieWithError(common.CreateDirectory(runDirectory, false))

	fmt.Println(config)
	return config
}

var AgentConfig = ParseArguments()
