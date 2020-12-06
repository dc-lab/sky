package helpers

import (
	cloud "github.com/dc-lab/sky/api/proto/cloud_manager"
	"github.com/dc-lab/sky/api/proto/common"
)

func NewAwsApiSettings(region string) cloud.TApiSettings {
	return cloud.TApiSettings{
		Body: &cloud.TApiSettings_AmazonSettings{
			AmazonSettings: &cloud.TAmazonApiSettings{
				Scheme: cloud.TAmazonApiSettings_HTTPS,
				Region: region,
			},
		},
	}
}

func NewAwsCredentials(accessKeyId, secretAccessKey string) cloud.TUserCredentials {
	return cloud.TUserCredentials{
		Body: &cloud.TUserCredentials_AmazonCredentials{
			AmazonCredentials: &cloud.TAmazonCredentials{
				AccessKeyId:     accessKeyId,
				SecretAccessKey: secretAccessKey,
			},
		},
	}
}

func NewHardwareData(coresCount float64, memoryBytes uint64, diskBytes uint64) common.THardwareData {
	return common.THardwareData{
		CoresCount:  coresCount,
		MemoryBytes: memoryBytes,
		DiskBytes:   diskBytes,
	}
}

func NewDockerImage(registry, repository, image, tag string) cloud.TDockerImage {
	return cloud.TDockerImage{
		Registry:   registry,
		Repository: repository,
		Image:      image,
		Tag:        tag,
	}
}
