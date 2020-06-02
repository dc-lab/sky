package cmd

import (
	"github.com/dc-lab/sky/cli/cmd/files"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "files operations",
}

func init() {
	RootCmd.AddCommand(filesCmd)
	filesCmd.AddCommand(files.CreateCmd)
}
