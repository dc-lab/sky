package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud_manager/client/helpers"
)

type StopCmdParams struct {
	instanceUuid string
}

var stopParams StopCmdParams

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "StopInstance",
	Long:  `Make GRPC call to Cloud server with StopInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_StopInstanceRequest{
				StopInstanceRequest: &cloud.TStopInstanceRequest{
					Creds:        &creds,
					Settings:     &apiSettings,
					InstanceUuid: stopParams.instanceUuid,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetStopInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("OK Instance stopped")
	},
}

func init() {
	stopCmd.Flags().StringVar(&stopParams.instanceUuid, "uuid", "", "Instance UUID")
	stopCmd.MarkPersistentFlagRequired("uuid")

	rootCmd.AddCommand(stopCmd)
}
