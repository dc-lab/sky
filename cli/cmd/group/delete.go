package group

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const deleteUrlSuffix = "/groups/"

var DeleteCmd = &cobra.Command{
	Use:   "delete [group id]",
	Short: "delete group",
	Long:  `Delete your group with specified id`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		groupId := args[0]
		if groupId == "" {
			log.Fatal("ID of group cannot be empty")
		}

		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodDelete, url + deleteUrlSuffix + groupId, nil, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
