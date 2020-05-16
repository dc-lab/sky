package app

import (
	"flag"
	"github.com/spf13/viper"
)

type config struct {
	HTTPAddress   string
	GRPCAddress   string
	LogsDir       string
	DBUser        string
	DBPasswordEnv string
	DBHost        string
	DBName        string
	DBSsl         bool
}

var Config = config{
	":8090",
	":5051",
	".",
	"oleg",
	"DB_PASSWORD",
	"rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	"sky_postgre",
	true,
}

func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to resource manager configuration")
	flag.Parse()

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
