package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	EnvAwsAccessKey   = "SKY_CLOUD_AWS_ACCESS_KEY"
	EnvAwsAccessKeyId = "SKY_CLOUD_AWS_ACCESS_KEY_ID"
)

func SetEnvDefaultForFlags(cmd *cobra.Command, envs map[string]string) {
	for env, flag := range envs {
		flag := cmd.Flags().Lookup(flag)
		updateFlag(flag, env)
	}
}

func SetEnvDefaultForPersistentFlags(cmd *cobra.Command, envs map[string]string) {
	for env, flag := range envs {
		flag := cmd.PersistentFlags().Lookup(flag)
		updateFlag(flag, env)
	}
}

func updateFlag(flag *pflag.Flag, env string) {
	flag.Usage = fmt.Sprintf("%v [env %v]", flag.Usage, env)
	if value, exist := os.LookupEnv(env); exist {
		flag.Value.Set(value)
	}
}
