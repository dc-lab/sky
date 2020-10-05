package cmd

import (
	"github.com/dc-lab/sky/internal/cli/cmd/file"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "file operations",
}

func init() {
	RootCmd.AddCommand(fileCmd)
	fileCmd.AddCommand(file.CreateCmd)
	fileCmd.AddCommand(file.DownloadCmd)
}
