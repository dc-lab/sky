package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dc-lab/sky/cloud/client/helpers"
	pb "github.com/dc-lab/sky/cloud/proto"
)

type DeployCmdParams struct {
	instanceUuid string
	registry string
	repository string
	image string
	tag string
}

var deployParams DeployCmdParams

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "DeployImage call",
	Long: `Make GRPC call to Cloud server with DeployImage action`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")

		creds := helpers.NewAwsCredentials(rootParams.accessKeyId, rootParams.secretAccessKey)
		apiSettings := helpers.NewAwsApiSettings(rootParams.region)
		dockerImage := helpers.NewDockerImage(deployParams.registry, deployParams.repository, deployParams.image, deployParams.tag)

		req := pb.TCloudRequest{
			Body: &pb.TCloudRequest_DeployImageRequest{
				DeployImageRequest: &pb.TDeployImageRequest{
					Creds: &creds,
					Settings: &apiSettings,
					InstanceUuid: deallocateParams.instanceUuid,
					DockerImage: &dockerImage,
				},
			},
		}

		resp := helpers.MakeCloudRequest(rootParams.grpcPort, &req)

		typedResp := resp.GetDeployImageResponse()
		if typedResp == nil {
			log.Fatalf("unexpected result type")
		}

		helpers.EnsureOkStatusCode(typedResp)

		log.Printf("OK Image deployed")
	},
}

func init() {
	deployCmd.Flags().StringVar(&deployParams.instanceUuid, "uuid", "", "Instance UUID")
	deployCmd.MarkPersistentFlagRequired("uuid")

	deployCmd.Flags().StringVar(&deployParams.instanceUuid, "registry-url", "", "Docker registry URL")
	deployCmd.MarkPersistentFlagRequired("registry-url")

	deployCmd.Flags().StringVar(&deployParams.instanceUuid, "repository", "", "Docker repository")
	deployCmd.MarkPersistentFlagRequired("repository")

	deployCmd.Flags().StringVar(&deployParams.instanceUuid, "image", "", "Docker image")

	deployCmd.Flags().StringVar(&deployParams.instanceUuid, "tag", "", "Docker image tag")

	rootCmd.AddCommand(deployCmd)
}
