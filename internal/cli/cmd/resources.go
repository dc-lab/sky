package cmd

import (
	"github.com/dc-lab/sky/internal/cli/cmd/resources"
	"github.com/spf13/cobra"
)

var resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "resources operations",
	Long:  `Create, delete and get your resources`,
}

func init() {
	RootCmd.AddCommand(resourcesCmd)
	resourcesCmd.AddCommand(resources.CreateCmd)
	resourcesCmd.AddCommand(resources.DeleteCmd)
	resourcesCmd.AddCommand(resources.GetCmd)
	resourcesCmd.AddCommand(resources.ListCmd)
}
