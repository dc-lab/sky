package cmd

import (
	"github.com/dc-lab/sky/cli/utils"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

const registerUrlSuffix = "/register"

var registerCmd = &cobra.Command{
	Use:   "register [login] [password]",
	Short: "register in sky",
	Long:  `Register in sky platform`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)

		login := args[0]
		if login == "" {
			log.Fatal("Login cannot be empty")
		}

		password := args[1]
		if password == "" {
			log.Fatal("Password cannot be empty")
		}

		req := map[string]string{"login": login, "password": password}
		statusCode, body := utils.MakeRequest(http.MethodPost, url + registerUrlSuffix, &req, nil)

		log.Println(statusCode)
		log.Println(body)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
}
