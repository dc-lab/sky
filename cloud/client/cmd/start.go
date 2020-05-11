package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud/client/helpers"
)

type StartCmdParams struct {
	instanceUuid string
}

var startParams StartCmdParams

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "StartInstance",
	Long:  `Make GRPC call to Cloud server with StartInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_StartInstanceRequest{
				StartInstanceRequest: &cloud.TStartInstanceRequest{
					Creds:        &creds,
					Settings:     &apiSettings,
					InstanceUuid: startParams.instanceUuid,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetStartInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("OK Instance started")
	},
}

func init() {
	startCmd.Flags().StringVar(&startParams.instanceUuid, "uuid", "", "Instance UUID")
	startCmd.MarkPersistentFlagRequired("uuid")

	rootCmd.AddCommand(startCmd)
}
