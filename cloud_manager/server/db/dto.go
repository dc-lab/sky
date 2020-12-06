package db

import (
	"github.com/dc-lab/sky/cloud_manager/server/entity"
)

type Credentials struct {
	Id             string
	OwnerId        string
	DisplayName    string
	Provider       entity.CloudProvider
	AwsAccessKey   *string
	AwsAccessKeyId *string
}

type ResourceSpec struct {
	CpuCores    float64
	MemoryBytes uint64
	DiskBytes   uint64
}

type CloudResourceFactory struct {
	Id                       string
	OwnerId                  string
	Token                    string
	DisplayName              string
	Provider                 entity.CloudProvider
	AgentDockerVersion       string
	Type                     entity.ResourceType
	InstanceCpuLimitCores    float64
	InstanceMemoryLimitBytes uint64
	InstanceDiskLimitBytes   uint64
	FactoryCpuLimitCores     float64
	FactoryMemoryLimitBytes  uint64
	FactoryDiskLimitBytes    uint64
	UsedCpuCores             float64
	UsedMemoryBytes          uint64
	UsedDiskBytes            uint64
	InstanceLimitCount       uint64
	UsedInstanceCount        uint64
	AwsCluster               *string
	AwsVpc                   *string
	AwsTag                   *string
	AwsNets                  *[]string
	AwsSecGroups             *[]string
}

type CloudResource struct {
	Id                   string
	FactoryId            string
	Token                string
	DisplayName          string
	Status               entity.ResourceStatus
	CpuLimitCores        float64
	MemoryLimitBytes     uint64
	DiskLimitBytes       uint64
	CpuGuaranteeCores    float64
	MemoryGuaranteeBytes uint64
	DiskGuaranteeBytes   uint64
}

type Transaction struct {
	Id         string
	ExternalId *string
	ExpireAt   int64
	Status     entity.TransactionStatus
}
