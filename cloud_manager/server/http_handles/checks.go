package http_handles

import (
	"fmt"
	"github.com/dc-lab/sky/cloud_manager/server/entity"
)

func validateCredentialsInput(ci *entity.CreateCredentialsInput) error {
	if ci.Provider == entity.ProviderAWS {
		if ci.AwsAccessKeyId == nil {
			return fmt.Errorf("invalid AWS credentials format: access key id not provided")
		}
		if ci.AwsAccessKey == nil {
			return fmt.Errorf("invalid AWS credentials format: access key not provided")
		}
		return nil
	}
	return fmt.Errorf("unrecognized cloud provided: " + ci.Provider.String())
}
