package user

import (
	"github.com/dc-lab/sky/internal/cli/utils"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const changePasswordUrlSuffix = "/change_password"

var ChangePasswordCmd = &cobra.Command{
	Use:   "change_password [old_password] [new_password]",
	Short: "change password for current account",
	Long:  `Change password for current account`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		oldPassword := args[0]
		if oldPassword == "" {
			log.Fatal("Empty old password")
		}
		newPassword := args[1]
		if newPassword == "" {
			log.Fatal("Empty new password")
		}

		req := map[string]string{"old_password": oldPassword, "new_password": newPassword}
		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodPost, url+changePasswordUrlSuffix, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
