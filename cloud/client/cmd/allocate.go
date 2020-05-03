package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud/client/helpers"
)

type AllocateCmdParams struct {
	coresCount uint32
	memoryBytes uint64
	diskBytes uint64
}

var allocateParams AllocateCmdParams

// allocateCmd represents the allocate command
var allocateCmd = &cobra.Command{
	Use:   "allocate",
	Short: "AllocateInstance call",
	Long: `Make GRPC call to Cloud server with AllocateInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("allocate called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)
		hardwareData := helpers.NewHardwareData(allocateParams.coresCount, allocateParams.memoryBytes, allocateParams.diskBytes)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_AllocateInstanceRequest{
				AllocateInstanceRequest: &cloud.TAllocateInstanceRequest{
					Creds: &creds,
					Settings: &apiSettings,
					HardwareData: &hardwareData,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetAllocateInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("instance_uuid: %s", typedResp.InstanceUuid)
	},
}

func init() {
	allocateCmd.Flags().Uint32Var(&allocateParams.coresCount, "cpu", 0, "Cores count")

	allocateCmd.Flags().Uint64Var(&allocateParams.memoryBytes, "memory", 0, "RAM in bytes")

	allocateCmd.Flags().Uint64Var(&allocateParams.diskBytes, "hdd", 0, "HDD in bytes")

	rootCmd.AddCommand(allocateCmd)
}
