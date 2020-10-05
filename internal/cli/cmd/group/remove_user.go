package group

import (
	"github.com/dc-lab/sky/internal/cli/utils"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const removeUserUrlSuffix = "/groups/"

var RemoveUserCmd = &cobra.Command{
	Use:   "remove_user [group id] [user id]",
	Short: "remove user from group",
	Long:  `Remove user from specific group`,
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

		req := map[string][]string{"users_to_add": {}, "users_to_remove": {userId}}
		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeOtherRequest(http.MethodPost, url + removeUserUrlSuffix + groupId, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
