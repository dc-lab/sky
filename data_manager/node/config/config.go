package config

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	HttpBindAddress string `mapstructure:"http_bind_address"`
	AccessAddress   string `mapstructure:"access_address"`
	LogFile         string `mapstructure:"log_file"`

	MasterAddress string        `mapstructure:"master_address"`
	PushInterval  time.Duration `mapstructure:"push_interval"`

	StorageDir  string `mapstructure:"storage_dir"`
	MaxFileSize int64  `mapstructure:"max_file_size"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("data_manager_node")
	viper.SetEnvPrefix("dm_node")
	viper.AddConfigPath(".")

	viper.BindEnv("HTTP_BIND_ADDRESS")
	viper.BindEnv("ACCESS_ADDRESS")
	viper.BindEnv("LOG_FILE")
	viper.BindEnv("MASTER_ADDRESS")
	viper.BindEnv("PUSH_INTERVAL")
	viper.BindEnv("STORAGE_DIR")
	viper.BindEnv("MAX_FILE_SIZE")

	viper.SetDefault("max_file_size", 1*1024*1024*1024) // 1GiB
	viper.SetDefault("push_interval", 10*time.Second)

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
