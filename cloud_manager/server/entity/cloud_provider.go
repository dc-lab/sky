package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CloudProvider int

const (
	ProviderUnknown CloudProvider = iota
	ProviderAWS
)

func (cp CloudProvider) String() string {
	return providerToString[cp]
}

var providerToString = map[CloudProvider]string{
	ProviderUnknown: "unknown",
	ProviderAWS:     "aws",
}

var providerToID = map[string]CloudProvider{
	"aws": ProviderAWS,
}

func (cp CloudProvider) MarshalJSON() ([]byte, error) {
	if cp == ProviderUnknown {
		return nil, fmt.Errorf("cannot serialize unknown cloud provider")
	}
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(providerToString[cp])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (cp *CloudProvider) UnmarshalJSON(b []byte) error {
	var idx string
	err := json.Unmarshal(b, &idx)
	if err != nil {
		return err
	}
	// NOTE: return Unknown for unrecognized id
	*cp = providerToID[idx]
	return nil
}
