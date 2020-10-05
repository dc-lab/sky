package gateway

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	BindAddress        string `mapstructure:"bind_address"`
	LogFile            string `mapstructure:"log_file"`
	JobManagerAddress  string `mapstructure:"jm_address"`
	DataManagerAddress string `mapstructure:"dm_address"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("jmm")
	viper.SetEnvPrefix("JM")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	pflag.String("log_file", "", "Log file path")
	pflag.String("bind_address", "", "Http bind address")
	pflag.String("jm_address", "", "Job manager address")
	pflag.String("dm_address", "", "Data manager address")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn("Config file not found")
		} else {
			return nil, err
		}
	}

	// Viper does not fill env from AutomaticEnv() in Unmarshal()
	// See https://github.com/spf13/viper/issues/188
	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
