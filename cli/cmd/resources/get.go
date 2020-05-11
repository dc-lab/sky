package resources

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const getUrlSuffix = "/resources/"

var GetCmd = &cobra.Command{
	Use:   "get [resource id]",
	Short: "get resource info",
	Long:  `Get info about specific resource`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		userId, _ := cmd.Flags().GetString("user_id")
		resourceId := args[0]
		if resourceId == "" {
			fmt.Println("ID of resource cannot be empty")
			os.Exit(1)
		}
		request, err := http.NewRequest(http.MethodGet, url + getUrlSuffix + resourceId, nil)
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
