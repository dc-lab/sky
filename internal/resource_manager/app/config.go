package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	HttpBindAddress string `mapstructure:"http_bind_address"`
	GrpcBindAddress string `mapstructure:"grpc_bind_address"`
	LogFile         string `mapstructure:"log_file"`
	PostgresAddress string `mapstructure:"postgres_address"`

	DataManagerAddress string `mapstructure:"dm_address"`
}

var Config = config{}

func LoadConfig() error {
	viper.SetConfigName("mm")
	viper.SetEnvPrefix("RM")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	pflag.String("log_file", "", "Log file path")
	pflag.String("http_bind_address", "", "Http bind address")
	pflag.String("grpc_bind_address", "", "Grpc bind address")
	pflag.String("dm_address", "", "Data manager address")
	pflag.String("postgres_address", "", "Database address")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("Config file not found")
		} else {
			return err
		}
	}

	// Viper does not fill env from AutomaticEnv() in Unmarshal()
	// See https://github.com/spf13/viper/issues/188
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}
