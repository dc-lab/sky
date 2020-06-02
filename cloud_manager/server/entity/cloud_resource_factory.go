package entity

// TODO: check json parser for uint64 and float64
type CloudResourceFactory struct {
	Id                    string        `json:"id"`
	OwnerId 		      string        `json:"owner_id"`
	Token                 string        `json:"token"`
	DisplayName           string        `json:"display_name"`
	Provider              CloudProvider `json:"provider"`
	AgentDockerVersion    string        `json:"agent_docker_version"`
	Type                  ResourceType  `json:"type"`
	InstanceLimit         ResourceSpec  `json:"instance_limit"`
	FactoryLimit          ResourceSpec  `json:"factory_limit"`
	Used                  ResourceSpec  `json:"used"`
	InstanceCountLimit    uint64        `json:"instance_count_limit"`
	UsedInstanceCount     uint64        `json:"used_instance_count"`
	AwsCluster            *string        `json:"aws_cluster,omitempty"`
	AwsVpc                *string        `json:"aws_vpc,omitempty"`
	AwsTag                *string        `json:"aws_tag,omitempty"`
	AwsNets               *[]string      `json:"aws_nets,omitempty"`
	AwsSecGroups          *[]string      `json:"aws_sec_groups,omitempty"`
}
