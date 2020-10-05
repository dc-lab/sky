package file

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/internal/cli/utils"
	dm_client "github.com/dc-lab/sky/internal/data_manager/client"
)

var CreateCmd = &cobra.Command{
	Use:   "create --name <NAME> --file <PATH>",
	Short: "create file",
	Long: `Create file
Returns created file id`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		token := utils.GetUserToken(cmd)
		name := utils.GetVariable(cmd, "name")
		file := utils.GetVariable(cmd, "file")

		var reader *os.File
		if file == "-" {
			reader = os.Stdin
		} else {
			var err error
			reader, err = os.Open(file)

			if err != nil {
				log.Fatalf("Failed to open file: %s", err.Error())
			}
		}
		defer reader.Close()

		client, err := dm_client.MakeClient(url, token)
		if err != nil {
			log.Fatalf("Failed to connect to the platform: %s", err.Error())
		}

		res, err := client.CreateFileWithContents(&dm_client.File{
			Name: name,
		}, reader)
		if err != nil {
			log.Fatalf("Failed to upload file: %s", err.Error())
		}

		fmt.Printf("Successfully uploaded file")
		fmt.Printf("ID: %s", res.Id)
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "n", "", "Name of file")
	CreateCmd.MarkFlagRequired("name")

	CreateCmd.Flags().StringP("file", "f", "", "File to upload")
	CreateCmd.MarkFlagRequired("file")
}
