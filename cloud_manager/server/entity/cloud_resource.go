package entity

import "github.com/dc-lab/sky/cloud_manager/server/db"

type CloudResource struct {
	Id        string         `json:"id"`
	FactoryId string         `json:"factory_id"`
	Token     string         `json:"token"`
	DisplayName string       `json:"display_name"`
	Status    ResourceStatus `json:"status"`
	Limit     ResourceSpec   `json:"limit"`
	Guarantee ResourceSpec   `json:"guarantee"`
}

func CloudResourceFromDto(resource db.CloudResource) CloudResource {
	return CloudResource{
		Id:          resource.Id,
		FactoryId:   resource.FactoryId,
		Token:       resource.Token,
		DisplayName: resource.DisplayName,
		Status:      resource.Status,
		Limit:       ResourceSpec{
			CpuCores:    resource.CpuLimitCores,
			MemoryBytes: resource.MemoryLimitBytes,
			DiskBytes:   resource.DiskLimitBytes,
		},
		Guarantee:   ResourceSpec{
			CpuCores:    resource.CpuGuaranteeCores,
			MemoryBytes: resource.MemoryGuaranteeBytes,
			DiskBytes:   resource.DiskGuaranteeBytes,
		},
	}
}
