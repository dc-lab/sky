package resources

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const listUrlSuffix = "/resources"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list resources",
	Long:  `List all available resources`,
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
