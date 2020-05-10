package resources

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const listUrlSuffix = "/resources"

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list resources",
	Long:  `List all available resources`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		userId, _ := cmd.Flags().GetString("user_id")
		request, err := http.NewRequest(http.MethodGet, url + listUrlSuffix, nil)
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
