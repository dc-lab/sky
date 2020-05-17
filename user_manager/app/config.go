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
	HTTPAddress:   ":6272",
	GRPCAddress:   ":6273",
	LogsDir:       ".",
	DBUser:        "oleg",
	DBPasswordEnv: "DB_PASSWORD",
	DBHost:        "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	DBName:        "sky_postgre",
	DBSsl:         true,
}

func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to user manager configuration")
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
