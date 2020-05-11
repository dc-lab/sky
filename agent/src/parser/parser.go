package parser

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	common "github.com/dc-lab/sky/agent/src/common"
	"github.com/spf13/viper"
)

type Config struct {
	ResourceManagerAddress string
	AgentDirectory         string
	Token                  string
}

func GetToken(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		common.DealWithError(err)
		os.Exit(1)
	}
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
	var configPath = *flag.String("config", "config.json", "Path to agent configuration file")
	flag.Parse()
	v, err := readConfig(configPath, map[string]interface{}{
		"ResourceManagerAddress": "localhost:5051",
		"AgentDirectory":         "/tmp/agent",
		"TokenPath":              "/tmp/token",
	})
	if err != nil {
		common.DealWithError(err)
		os.Exit(1)
	}
	token := GetToken(v.GetString("TokenPath"))
	config := Config{
		ResourceManagerAddress: v.GetString("ResourceManagerAddress"),
		AgentDirectory:         v.GetString("AgentDirectory"),
		Token:                  token,
	}
	if flag, err := common.PathExists(config.AgentDirectory, false); flag {
		common.DealWithError(err)
		os.RemoveAll(config.AgentDirectory)
	}
	err = os.Mkdir(config.AgentDirectory, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	return config
}

var AgentConfig = ParseArguments()
