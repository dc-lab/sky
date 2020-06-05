package app

import (
	"flag"
	"regexp"

	"github.com/spf13/viper"
)

type Endpoint struct {
	PathPattern  string
	Hostname     string
	AuthOptional bool
	PathRegex    *regexp.Regexp
}

type config struct {
	HTTPAddress   string     `mapstructure:"http_address"`
	LogsDir       string     `mapstructure:"logs_dir"`
	Endpoints     []Endpoint `mapstructure:"endpoints"`
	DBUser        string     `mapstructure:"db_user"`
	DBPasswordEnv string
	DBHost        string `mapstructure:"db_host"`
	DBName        string `mapstructure:"db_name"`
	DBSsl         bool   `mapstructure:"db_ssl"`
}

var Config = config{
	HTTPAddress:   ":4000",
	LogsDir:       ".",
	Endpoints:     []Endpoint{},
	DBUser:        "oleg",
	DBPasswordEnv: "RP_DB_PASSWORD",
	DBHost:        "rc1b-6marivlovkr6pccx.mdb.yandexcloud.net:6432",
	DBName:        "sky_postgre",
	DBSsl:         true,
}

func ParseConfig() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "Path to reverse proxy configuration")
	flag.Parse()

	viper.SetConfigFile(configPath)

	// See https://github.com/spf13/viper/issues/188
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RP")
	viper.BindEnv("HTTP_ADDRESS")
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

	err = compileRegexps()
	if err != nil {
		panic(err)
	}
}

func compileRegexps() error {
	for i, endpoint := range Config.Endpoints {
		if endpoint.PathPattern[0] != '^' {
			endpoint.PathPattern = "^" + endpoint.PathPattern
		}
		regex, err := regexp.Compile(endpoint.PathPattern)
		if err != nil {
			return err
		}

		Config.Endpoints[i].PathRegex = regex
	}
	return nil
}
