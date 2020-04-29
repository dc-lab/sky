package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	GrpcPort uint16

	rootCmd = &cobra.Command{
		Use: "server.go [ARGS]",
		Short: "Run SKY cloud API",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Uint16Var(&GrpcPort, "grpc-port", 5005, "")
}
