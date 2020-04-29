package helpers

import pb "github.com/dc-lab/sky/cloud/proto"

func NewAwsApiSettings(region string) pb.TApiSettings {
	return pb.TApiSettings{
		Body: &pb.TApiSettings_AmazonSettings{
			AmazonSettings: &pb.TAmazonApiSettings{
				Scheme: pb.TAmazonApiSettings_HTTPS,
				Region: region,
			},
		},
	}
}

func NewAwsCredentials(accessKeyId, secretAccessKey string) pb.TUserCredentials {
	return pb.TUserCredentials{
		Body: &pb.TUserCredentials_AmazonCredentials{
			AmazonCredentials: &pb.TAmazonCredentials{
				AccessKeyId: accessKeyId,
				SecretAccessKey: secretAccessKey,
			},
		},
	}
}

func NewHardwareData(coresCount uint32, memoryBytes uint64, diskBytes uint64) pb.THardwareData {
	return pb.THardwareData{
		CoresCount: coresCount,
		MemoryBytes: memoryBytes,
		DiskBytes: diskBytes,
	}
}

func NewDockerImage(registry, repository, image, tag string) pb.TDockerImage {
	return pb.TDockerImage{
		Registry: registry,
		Repository: repository,
		Image: image,
		Tag: tag,
	}
}
