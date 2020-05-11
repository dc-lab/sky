package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/cloud_manager/client/config"
)

var (
	envs = map[string]string{
		config.EnvAwsAccessKey:   "aws-access-key",
		config.EnvAwsAccessKeyId: "aws-access-key-id",
	}

	rootParams RootCmdParams
)

type RootCmdParams struct {
	grpcPort        uint16
	accessKeyId     string
	secretAccessKey string
	region          string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "CLI client for Cloud GRPC server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint16Var(&rootParams.grpcPort, "grpc-port", 5005, "Cloud server port")

	rootCmd.PersistentFlags().StringVar(&rootParams.secretAccessKey, "aws-access-key", "", "AWS access key")

	rootCmd.PersistentFlags().StringVar(&rootParams.accessKeyId, "aws-access-key-id", "", "AWS access key id")

	rootCmd.PersistentFlags().StringVar(&rootParams.region, "aws-region", "", "AWS region for API requests")

	config.SetEnvDefaultForPersistentFlags(rootCmd, envs)
}
