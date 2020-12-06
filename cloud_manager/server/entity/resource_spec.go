package entity

type ResourceSpec struct {
	CpuCores    float64 `json:"cpu_cores"`
	MemoryBytes uint64  `json:"memory_bytes"`
	DiskBytes   uint64  `json:"disk_bytes"`
}
