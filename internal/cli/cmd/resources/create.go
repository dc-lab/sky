package resources

import (
	"github.com/dc-lab/sky/internal/cli/utils"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const createUrlSuffix = "/resources"

var CreateCmd = &cobra.Command{
	Use:   "create [name] [type]",
	Short: "create resource",
	Long: `Create resource with specified name and type
Returns resource object with all data you need
Type: single or pool
`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		userToken := utils.GetUserToken(cmd)
		resourceName := args[0]
		if resourceName == "" {
			log.Fatal("Resource name cannot be empty")
		}
		resourceType := args[1]
		if resourceType != "single" && resourceType != "pool" {
			log.Fatal("Resource type must be either 'single' or 'pool'")
		}

		req := map[string]string{"name": resourceName, "type": resourceType}
		headers := map[string]string{"User-Token": userToken}
		statusCode, body := utils.MakeRequest(http.MethodPost, url + createUrlSuffix, &req, &headers)

		log.Println(statusCode)
		log.Println(body)
	},
}
