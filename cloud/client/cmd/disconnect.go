package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud/client/helpers"
)

type DisconnectCmdParams struct {
	instanceUuid string
}

var disconnectParams DisconnectCmdParams

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "DisconnectInstance",
	Long:  `Make GRPC call to Cloud server with DisconnectInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("disconnect called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_DisconnectInstanceRequest{
				DisconnectInstanceRequest: &cloud.TDisconnectInstanceRequest{
					Creds:        &creds,
					InstanceUuid: disconnectParams.instanceUuid,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetDisconnectInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("OK Instance disconnected")
	},
}

func init() {
	disconnectCmd.Flags().StringVar(&disconnectParams.instanceUuid, "uuid", "", "Instance UUID")
	disconnectCmd.MarkPersistentFlagRequired("uuid")

	rootCmd.AddCommand(disconnectCmd)
}
