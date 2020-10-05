package user

import (
	"github.com/dc-lab/sky/internal/cli/utils"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const loginUrlSuffix = "/login"

var LoginCmd = &cobra.Command{
	Use:   "login [login] [password]",
	Short: "login in sky",
	Long:  `Login in sky platform`,
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
		statusCode, body := utils.MakeRequest(http.MethodPost, url+loginUrlSuffix, &req, nil)

		log.Println(statusCode)
		log.Println(body)
	},
}
