package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	BindAddress string `mapstructure:"bind_address"`
	LogFile     string `mapstructure:"log_file"`
	StorageDir  string `mapstructure:"storage_dir"`
	MaxFileSize int64  `mapstructure:"max_file_size"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("data_manager")
	viper.SetEnvPrefix("dm")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetDefault("bind_address", ":8080")
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
