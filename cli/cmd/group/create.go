package group

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const createUrlSuffix = "/groups"

var CreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "create group",
	Long: `Create group
Returns group object`,
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
		statusCode, body := utils.MakeRequest(http.MethodPost, url + createUrlSuffix, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}