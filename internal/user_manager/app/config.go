package app

import (
	"flag"
	"github.com/spf13/viper"
)

type config struct {
	HTTPAddress   string `mapstructure:"http_address"`
	GRPCAddress   string `mapstructure:"grpc_address"`
	LogsDir       string `mapstructure:"logs_dir"`
	DBUser        string `mapstructure:"db_user"`
	DBPasswordEnv string
	DBHost        string `mapstructure:"db_host"`
	DBName        string `mapstructure:"db_name"`
	DBSsl         bool   `mapstructure:"db_ssl"`
}

var Config = config{
	HTTPAddress:   ":6272",
	GRPCAddress:   ":6273",
	LogsDir:       ".",
	DBUser:        "oleg",
	DBPasswordEnv: "UM_DB_PASSWORD",
	DBHost:        "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	DBName:        "sky_postgre",
	DBSsl:         true,
}

func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to user manager configuration")
	flag.Parse()

	viper.SetConfigFile(configPath)

	// See https://github.com/spf13/viper/issues/188
	viper.AutomaticEnv()
	viper.SetEnvPrefix("UM")
	viper.BindEnv("HTTP_ADDRESS")
	viper.BindEnv("GRPC_ADDRESS")
	viper.BindEnv("LOGS_DIR")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSL")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
