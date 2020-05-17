package resources

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const getUrlSuffix = "/resources/"

var GetCmd = &cobra.Command{
	Use:   "get [resource id]",
	Short: "get resource info",
	Long:  `Get info about specific resource`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		resourceId := args[0]
		if resourceId == "" {
			log.Fatal("ID of resource cannot be empty")
		}

		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodGet, url + getUrlSuffix + resourceId, nil, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
