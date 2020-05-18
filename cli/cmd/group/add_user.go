package group

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const addUserUrlSuffix = "/groups/"

var AddUserCmd = &cobra.Command{
	Use:   "add_user [group id] [user id]",
	Short: "add user to group",
	Long:  `Add user to specific group`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		groupId := args[0]
		if groupId == "" {
			log.Fatal("ID of group cannot be empty")
		}
		userId := args[1]
		if userId == "" {
			log.Fatal("ID of user cannot be empty")
		}

		req := map[string][]string{"users_to_add": []string{userId}, "users_to_remove": []string{}}
		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeOtherRequest(http.MethodPost, url + addUserUrlSuffix + groupId, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
