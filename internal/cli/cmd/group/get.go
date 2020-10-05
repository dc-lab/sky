package group

import (
	"github.com/dc-lab/sky/internal/cli/utils"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const getUrlSuffix = "/groups/"

var GetCmd = &cobra.Command{
	Use:   "get [group id]",
	Short: "get group info",
	Long:  `Get info about specific group`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		groupId := args[0]
		if groupId == "" {
			log.Fatal("ID of group cannot be empty")
		}

		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodGet, url + getUrlSuffix + groupId, nil, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}

