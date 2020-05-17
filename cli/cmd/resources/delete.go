package resources

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const deleteUrlSuffix = "/resources/"

var DeleteCmd = &cobra.Command{
	Use:   "delete [resource id]",
	Short: "delete resource",
	Long:  `Delete your resource with specified id`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		resourceId := args[0]
		if resourceId == "" {
			log.Fatal("ID of resource cannot be empty")
		}

		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodDelete, url + deleteUrlSuffix + resourceId, nil, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
