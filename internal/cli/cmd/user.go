package cmd

import (
	"github.com/dc-lab/sky/internal/cli/cmd/user"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user operations",
	Long:  `Login, register or change password`,
}

func init() {
	RootCmd.AddCommand(userCmd)
	userCmd.AddCommand(user.LoginCmd)
	userCmd.AddCommand(user.RegisterCmd)
	userCmd.AddCommand(user.ChangePasswordCmd)
}
