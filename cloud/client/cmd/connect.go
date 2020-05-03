package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/api/proto/cloud"
	"github.com/dc-lab/sky/cloud/client/helpers"
)

type ConnectCmdParams struct {
	coresCount uint32
	memoryBytes uint64
	diskBytes uint64
}

var connectParams ConnectCmdParams

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "ConnectInstance call",
	Long: `Make GRPC call to Cloud server with ConnectInstance action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("connect called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)
		hardwareData := helpers.NewHardwareData(connectParams.coresCount, connectParams.memoryBytes, connectParams.diskBytes)

		req := cloud.TCloudRequest{
			Body: &cloud.TCloudRequest_ConnectInstanceRequest{
				ConnectInstanceRequest: &cloud.TConnectInstanceRequest{
					Creds: &creds,
					Settings: &apiSettings,
					HardwareData: &hardwareData,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetConnectInstanceResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("instance_uuid: %s", typedResp.InstanceUuid)
	},
}

func init() {
	connectCmd.Flags().Uint32Var(&connectParams.coresCount, "cpu", 0, "Cores count")

	connectCmd.Flags().Uint64Var(&connectParams.memoryBytes, "memory", 0, "RAM in bytes")

	connectCmd.Flags().Uint64Var(&connectParams.diskBytes, "hdd", 0, "HDD in bytes")

	rootCmd.AddCommand(connectCmd)
}
