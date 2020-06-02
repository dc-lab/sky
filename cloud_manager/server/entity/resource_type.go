package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ResourceType int

const (
	ResourceTypeUnknown ResourceType = iota
	ResourceTypeIaaS
	ResourceTypeFaaS
)

func (rt ResourceType) String() string {
	return resourceTypeToString[rt]
}

var resourceTypeToString = map[ResourceType]string{
	ResourceTypeUnknown: "unknown",
	ResourceTypeIaaS:    "iaas",
	ResourceTypeFaaS:    "faas",
}

var resourceTypeToID = map[string]ResourceType{
	"iaas": ResourceTypeIaaS,
	"faas": ResourceTypeFaaS,
}

func (rt ResourceType) MarshalJSON() ([]byte, error) {
	if rt == ResourceTypeUnknown {
		return nil, fmt.Errorf("cannot serialize unknown resource type")
	}
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(resourceTypeToString[rt])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (rt *ResourceType) UnmarshalJSON(b []byte) error {
	var idx string
	err := json.Unmarshal(b, &idx)
	if err != nil {
		return err
	}
	// NOTE: return Unknown for unrecognized id
	*rt = resourceTypeToID[idx]
	return nil
}
