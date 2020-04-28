package resources

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const deleteUrlSuffix = "/resources/"

var DeleteCmd = &cobra.Command{
	Use:   "delete [resource id]",
	Short: "delete resource",
	Long:  `Delete your resource with specified id`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		userId, _ := cmd.Flags().GetString("user_id")
		resourceId := args[0]
		if resourceId == "" {
			fmt.Println("ID of resource cannot be empty")
			os.Exit(1)
		}
		request, err := http.NewRequest(http.MethodDelete, url + deleteUrlSuffix + resourceId, nil)
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
