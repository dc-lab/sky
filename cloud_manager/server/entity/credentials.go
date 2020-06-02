package entity

import "github.com/dc-lab/sky/cloud_manager/server/db"

type Credentials struct {
	Id             string        `json:"id"`
	OwnerId        string        `json:"owner_id"`
	DisplayName    string        `json:"display_name"`
	Provider       CloudProvider `json:"provider"`
	AwsAccessKey   *string       `json:"aws_access_key,omitempty"`
	AwsAccessKeyId *string       `json:"aws_access_key_id,omitempty"`
}

type CreateCredentialsInput struct {
	DisplayName    string        `json:"display_name"`
	Provider       CloudProvider `json:"provider"`
	AwsAccessKey   *string       `json:"aws_access_key,omitempty"`
	AwsAccessKeyId *string       `json:"aws_access_key_id,omitempty"`
}

func CredentialsFromDto(credsDto db.Credentials) Credentials {
	return Credentials{
		Id:             credsDto.Id,
		OwnerId:        credsDto.OwnerId,
		DisplayName:    credsDto.DisplayName,
		Provider:       credsDto.Provider,
		AwsAccessKey:   credsDto.AwsAccessKey,
		AwsAccessKeyId: credsDto.AwsAccessKeyId,
	}
}
