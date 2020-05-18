package cmd

import (
	"github.com/dc-lab/sky/cli/cmd/group"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "group operations",
	Long:  `Add, list`,
}

func init() {
	RootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(group.CreateCmd)
	groupCmd.AddCommand(group.ListCmd)
	groupCmd.AddCommand(group.GetCmd)
	groupCmd.AddCommand(group.DeleteCmd)
	groupCmd.AddCommand(group.AddUserCmd)
	groupCmd.AddCommand(group.RemoveUserCmd)
}
