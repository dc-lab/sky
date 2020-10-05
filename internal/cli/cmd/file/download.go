package file

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/internal/cli/utils"
	dm_client "github.com/dc-lab/sky/internal/data_manager/client"
)

var DownloadCmd = &cobra.Command{
	Use:   "download -output <PATH>",
	Short: "download file",
	Long:  `Download file using its id`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := utils.GetSkyUrl(cmd)
		token := utils.GetUserToken(cmd)
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Fatalf("Failed to get output file path from command line: %e", err)
		}
		id := args[0]
		if id == "" {
			log.Fatal("File id cannot be empty")
		}

		var writer *os.File
		if output == "" || output == "-" {
			writer = os.Stdout
		} else {
			var err error
			writer, err = os.Create(output)

			if err != nil {
				log.Fatalf("Failed to open file: %s", err.Error())
			}
		}
		defer writer.Close()

		client, err := dm_client.MakeClient(url, token)
		if err != nil {
			log.Fatalf("Failed to connect to the platform: %s", err.Error())
		}

		err = client.GetFileContents(id, writer)
		if err != nil {
			log.Fatalf("Failed to download file: %s", err.Error())
		}

		fmt.Printf("Successfully downloaded file")
	},
}

func init() {
	DownloadCmd.Flags().StringP("output", "o", "", "Name of the output file")
}
