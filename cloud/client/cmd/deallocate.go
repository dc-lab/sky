package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud/client/helpers"
)

type DeallocateCmdParams struct {
	instanceUuid string
}

var deallocateParams DeallocateCmdParams

// deallocateCmd represents the deallocate command
var deallocateCmd = &cobra.Command{
	Use:   "deallocate",
	Short: "DeallocateInstance call",
	Long: `Make GRPC call to Cloud server with DeallocateInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deallocate called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_DeallocateInstanceRequest{
				DeallocateInstanceRequest: &cloud.TDeallocateInstanceRequest{
					Creds: &creds,
					Settings: &apiSettings,
					InstanceUuid: deallocateParams.instanceUuid,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetDeallocateInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("OK Instance deallocated")
	},
}

func init() {
	deallocateCmd.Flags().StringVar(&deallocateParams.instanceUuid, "uuid", "", "Instance UUID")
	deallocateCmd.MarkPersistentFlagRequired("uuid")

	rootCmd.AddCommand(deallocateCmd)
}
