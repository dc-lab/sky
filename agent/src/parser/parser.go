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
	LocalCacheDirectory    string
	AgentLogFile           string
	HealthFile             string
	Token                  string
}

type CmdOptions struct {
	ConfigPath string
}

func GetToken(path string) string {
	bytes, err := ioutil.ReadFile(path)
	common.DieWithError(err)
	return string(bytes)
}

func ReadConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	if val, err := common.PathExists(filename, false); val && err == nil {
		v.SetConfigType("json")
		v.SetConfigName(filename)
		v.AddConfigPath(".")
	}
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

var AgentConfig Config

func ParseArguments() CmdOptions {
	var options CmdOptions
	flag.StringVar(&options.ConfigPath, "config", "config.json", "Path to agent configuration file")
	flag.Parse()
	return options
}

func InitializeAgentConfigFromOptions(options *CmdOptions) {
	viperObject, err := ReadConfig(options.ConfigPath, map[string]interface{}{
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/var/tmp/agent",
		"LogsDirectory":          "/var/tmp/agent-logs",
		"RunDirectory":           "/var/run/agent",
		"TokenPath":              "/var/tmp/token",
	})
	common.DieWithError(err)
	InitializeAgentConfig(viperObject)
}

func InitializeAgentConfigFromCustomFields(customFields map[string]interface{}) {
	viperObject, err := ReadConfig("", customFields)
	common.DieWithError(err)
	InitializeAgentConfig(viperObject)
}

func InitializeAgentConfig(v *viper.Viper) {
	token := GetToken(v.GetString("TokenPath"))
	logsDirectory := v.GetString("LogsDirectory")
	runDirectory := v.GetString("RunDirectory")
	config := Config{
		ResourceManagerAddress: v.GetString("ResourceManagerAddress"),
		AgentDirectory:         v.GetString("AgentDirectory"),
		AgentLogFile:           path.Join(logsDirectory, "agent.log"),
		HealthFile:             path.Join(runDirectory, "health.info"),
		Token:                  token,
		LocalCacheDirectory:    path.Join(v.GetString("AgentDirectory"), "local_cache"),
	}

	common.DieWithError(common.CreateDirectory(config.AgentDirectory, false))
	common.DieWithError(common.CreateDirectory(logsDirectory, false))
	common.DieWithError(common.CreateDirectory(runDirectory, false))
	common.DieWithError(common.CreateDirectory(config.LocalCacheDirectory, false))

	fmt.Println(config)
	AgentConfig = config
}
