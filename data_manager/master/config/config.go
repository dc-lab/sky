package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	GrpcBindAddress string `mapstructure:"grpc_bind_address"`
	HttpBindAddress string `mapstructure:"http_bind_address"`

	LogFile         string `mapstructure:"log_file"`
	PostgresAddress string `mapstructure:"postgres_address"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("data_manager")
	viper.SetEnvPrefix("dm_master")
	viper.AddConfigPath(".")

	viper.BindEnv("BIND_ADDRESS")
	viper.BindEnv("LOG_FILE")
	viper.BindEnv("STORAGE_DIR")
	viper.BindEnv("MAX_FILE_SIZE")
	viper.BindEnv("POSTGRES_ADDRESS")

	viper.SetDefault("http_bind_address", ":8080")
	viper.SetDefault("grpc_bind_address", ":8081")
	viper.SetDefault("max_file_size", 1*1024*1024*1024) // 1GiB

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("Config file not found")
		} else {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
