package app

import (
	"flag"
	"github.com/spf13/viper"
)

type Endpoint struct {
	PathPrefix   string
	Hostname     string
	AuthOptional bool
}

type config struct {
	HTTPAddress   string
	LogsDir       string
	Endpoints     []Endpoint
	DBUser        string
	DBPasswordEnv string
	DBHost        string
	DBName        string
	DBSsl         bool
}

var Config = config{
	HTTPAddress:   ":4000",
	LogsDir:       ".",
	Endpoints:     []Endpoint{},
	DBUser:        "oleg",
	DBPasswordEnv: "DB_PASSWORD",
	DBHost:        "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	DBName:        "sky_postgre",
	DBSsl:         true,
}

func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to reverse proxy configuration")
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
