package app

import (
	"flag"
	"path"

	"github.com/spf13/viper"

	"github.com/dc-lab/sky/cloud_manager/server/common"

)

type config struct {
	GRPCAddress   string
	LogsDir       string
	LogsFile      string
	DBUser        string
	DBPasswordEnv string
	DBHost        string
	DBName        string
	DBSsl         bool
}

var Config = config{
	GRPCAddress:   ":5151",
	LogsDir:       "/var/tmp/cloud-manager",
	DBUser:        "oleg",
	DBPasswordEnv: "DB_PASSWORD",
	DBHost:        "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	DBName:        "sky_postgre",
	DBSsl:         true,
}

// TODO: unify configs
func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to cloud manager configuration file")
	flag.Parse()

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	common.DieWithError("Cannot read config", err)

	err = viper.Unmarshal(&Config)
	common.DieWithError("Cannot unmarshal config", err)

	err = common.MakeDir(Config.LogsDir, false)
	common.DieWithError("Cannot make dir: " + Config.LogsDir, err)

	if Config.LogsFile == "" {
		Config.LogsFile = path.Join(Config.LogsDir, "cm.log")
	}
}
