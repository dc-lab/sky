package hardware

import (
	common "github.com/dc-lab/sky/internal/agent/src/common"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/struCoder/pidusage"
)

type HardwareData struct {
	CpuCount    float64
	MemoryBytes uint64
	DiskBytes   uint64
}

//////////////////BEGIN CORES//////////////////

func GetTotalCoresCount() float64 {
	corescount, err := cpu.Counts(false)
	common.DealWithError(err)
	return float64(corescount)
}

func GetFreeCoresCount() float64 {
	percentage, err := cpu.Percent(0, false)
	common.DealWithError(err)
	return GetTotalCoresCount() * (100. - percentage[0]) / 100.
}

//////////////////END CORES//////////////////

//////////////////BEGIN MEMORY//////////////////

func GetMemoryStat() *mem.VirtualMemoryStat {
	memorystat, err := mem.VirtualMemory()
	common.DealWithError(err)
	return memorystat
}

func GetTotalMemory() uint64 {
	return GetMemoryStat().Total
}

func GetFreeMemory() uint64 {
	return GetMemoryStat().Free
}

//////////////////END MEMORY//////////////////

//////////////////BEGIN DISK////////////////

func GetDiskStat(path string) *disk.UsageStat {
	diskamount, err := disk.Usage(path)
	common.DealWithError(err)
	return diskamount
}

func GetTotalDisk(path string) uint64 {
	return GetDiskStat(path).Total
}

func GetFreeDisk(path string) uint64 {
	return GetDiskStat(path).Free
}

func GetDiskUsage(path string) uint64 {
	return GetDiskStat(path).Used
}

//////////////////BEGIN DISK////////////////

func GetTotalHardwareData(agentDir string) HardwareData {
	return HardwareData{
		CpuCount:    GetTotalCoresCount(),
		MemoryBytes: GetTotalMemory(),
		DiskBytes:   GetTotalDisk(agentDir)}
}

func GetFreeHardwareData(agentDir string) HardwareData {
	return HardwareData{
		CpuCount:    GetFreeCoresCount(),
		MemoryBytes: GetFreeMemory(),
		DiskBytes:   GetFreeDisk(agentDir)}
}

func GetTaskHardwareDataUsage(pid int, taskPath string) HardwareData {
	sysInfo, err := pidusage.GetStat(pid)
	common.DealWithError(err)
	return HardwareData{
		CpuCount:    sysInfo.CPU,
		MemoryBytes: uint64(sysInfo.Memory),
		DiskBytes:   GetDiskStat(taskPath).Used,
	}
}
