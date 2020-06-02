package group

import (
	"github.com/spf13/cobra"
	"log"
	"net/http"

	"github.com/dc-lab/sky/cli/utils"
	dm_client "github.com/dc-lab/sky/data_manager/client"
)

const createUrlSuffix = "/groups"

var CreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "create file",
	Long: `Create file
Returns file object`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)

		groupName := args[0]
		if groupName == "" {
			log.Fatal("Group name cannot be empty")
		}

		req := map[string]string{"name": groupName}
		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodPost, url+createUrlSuffix, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
