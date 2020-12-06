package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ResourceStatus int

const (
	ResourceStatusUnknown ResourceStatus = iota
	ResourceStatusReserved
	ResourceStatusAllocating
	ResourceStatusAllocated
	ResourceStatusPreparing
	ResourceStatusPrepared
	ResourceStatusStarting
	ResourceStatusRunning
	ResourceStatusStopping
	ResourceStatusStopped
	ResourceStatusShuttingDown
	ResourceStatusDeallocating
)

func (rs ResourceStatus) String() string {
	return resourceStatusToString[rs]
}

var resourceStatusToString = map[ResourceStatus]string{
	ResourceStatusUnknown:      "unknown",
	ResourceStatusReserved:     "reserved",
	ResourceStatusAllocating:   "allocating",
	ResourceStatusAllocated:    "allocated",
	ResourceStatusPreparing:    "preparing",
	ResourceStatusPrepared:     "prepared",
	ResourceStatusStarting:     "starting",
	ResourceStatusRunning:      "running",
	ResourceStatusStopping:     "stopping",
	ResourceStatusStopped:      "stopped",
	ResourceStatusShuttingDown: "shutting_down",
	ResourceStatusDeallocating: "deallocating",
}

var resourceStatusToID = map[string]ResourceStatus{
	"unknown":       ResourceStatusUnknown,
	"reserved":      ResourceStatusReserved,
	"allocating":    ResourceStatusAllocating,
	"allocated":     ResourceStatusAllocated,
	"preparing":     ResourceStatusPreparing,
	"prepared":      ResourceStatusPrepared,
	"starting":      ResourceStatusStarting,
	"running":       ResourceStatusRunning,
	"stopping":      ResourceStatusStopping,
	"stopped":       ResourceStatusStopped,
	"shutting_down": ResourceStatusShuttingDown,
	"deallocating":  ResourceStatusDeallocating,
}

func (rs ResourceStatus) MarshalJSON() ([]byte, error) {
	if rs == ResourceStatusUnknown {
		return nil, fmt.Errorf("cannot serialize unknown resource status")
	}
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(resourceStatusToString[rs])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (rs *ResourceStatus) UnmarshalJSON(b []byte) error {
	var idx string
	err := json.Unmarshal(b, &idx)
	if err != nil {
		return err
	}
	// NOTE: return Unknown for unrecognized id
	*rs = resourceStatusToID[idx]
	return nil
}
