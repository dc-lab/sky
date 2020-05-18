package group

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const listUrlSuffix = "/groups"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list groups",
	Long:  `List all my groups`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodGet, url + listUrlSuffix, nil, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
