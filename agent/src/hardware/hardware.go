package hardware

import (
	common "github.com/dc-lab/sky/agent/src/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func GetCoresCount() uint32 {
	corescount, err := cpu.Counts(false)
	common.DealWithError(err)
	return uint32(corescount)
}

func GetMemory() uint64 {
	memorystat, err := mem.VirtualMemory()
	common.DealWithError(err)
	return memorystat.Total
}

func GetDisk() uint64 {
	diskamount, err := disk.Usage("/")
	common.DealWithError(err)
	return diskamount.Total
}

type HardwareData struct {
	CpuCount    uint32
	MemoryBytes uint64
	DiskBytes   uint64
}

func GetHardwareData() HardwareData {
	return HardwareData{
		CpuCount:    GetCoresCount(),
		MemoryBytes: GetMemory(),
		DiskBytes:   GetDisk()}
}
