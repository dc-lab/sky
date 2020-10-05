package utils

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetVariable(cmd *cobra.Command, variableName string) string {
	configValue := viper.GetString(variableName)
	if configValue != "" {
		return configValue
	}
	cmdValue, err := cmd.Flags().GetString(variableName)
	if err != nil {
		log.Printf("Error while getting '%s' value from command flags: %s\n", variableName, err)
		os.Exit(1)
	}
	return cmdValue
}

func GetSkyUrl(cmd *cobra.Command) string {
	return GetVariable(cmd, "url")
}

func GetUserToken(cmd *cobra.Command) string {
	return GetVariable(cmd, "token")
}
