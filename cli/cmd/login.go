package cmd

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

const loginUrlSuffix = "/login"

var loginCmd = &cobra.Command{
	Use:   "login [login] [password]",
	Short: "login in sky",
	Long:  `Login in sky platform`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		login := args[0]
		password := args[1]
		if login == "" {
			fmt.Println("Login cannot be empty")
			os.Exit(1)
		}
		if password == "" {
			fmt.Println("Password cannot be empty")
			os.Exit(1)
		}
		req := map[string]string{"login": login, "password": password}
		jsonRequest, _ := json.Marshal(req)

		request, err := http.NewRequest(http.MethodPost, url + loginUrlSuffix, bytes.NewBuffer(jsonRequest))
		if err != nil {
			log.Fatal(err)
		}
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
	RootCmd.AddCommand(loginCmd)
}
