package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		url, _ := cmd.Flags().GetString("url")
		userId, _ := cmd.Flags().GetString("user_id")
		resourceName := args[0]
		if resourceName == "" {
			fmt.Println("Resource name cannot be empty")
			os.Exit(1)
		}
		resourceType := args[1]
		if resourceType != "single" && resourceType != "pool" {
			fmt.Println("Resource type must be either 'single' or 'pool'")
			os.Exit(1)
		}
		req := map[string]string{"name": resourceName, "type": resourceType}
		jsonRequest, _ := json.Marshal(req)

		request, err := http.NewRequest(http.MethodPost, url + createUrlSuffix, bytes.NewBuffer(jsonRequest))
		if err != nil {
			log.Fatal(err)
		}
		request.Header.Set("User-Id", userId)
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			log.Printf("Something went wrong: %s", err)
			os.Exit(1)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		println(string(body))
	},
}

func init() {
}
