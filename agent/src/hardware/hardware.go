package hardware

import (
	common "github.com/dc-lab/sky/agent/src/common"
	"github.com/dc-lab/sky/agent/src/parser"
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

func GetTotalDisk() uint64 {
	return GetDiskStat(parser.AgentConfig.AgentDirectory).Total
}

func GetFreeDisk() uint64 {
	return GetDiskStat(parser.AgentConfig.AgentDirectory).Free
}

//////////////////BEGIN DISK////////////////

func GetTotalHardwareData() HardwareData {
	return HardwareData{
		CpuCount:    GetTotalCoresCount(),
		MemoryBytes: GetTotalMemory(),
		DiskBytes:   GetTotalDisk()}
}

func GetFreeHardwareData() HardwareData {
	return HardwareData{
		CpuCount:    GetFreeCoresCount(),
		MemoryBytes: GetFreeMemory(),
		DiskBytes:   GetFreeDisk()}
}

func GetTaskHardwareDataUsage(taskId string, pid int) HardwareData {
	sysInfo, err := pidusage.GetStat(pid)
	common.DealWithError(err)
	return HardwareData{
		CpuCount:    sysInfo.CPU,
		MemoryBytes: uint64(sysInfo.Memory),
		DiskBytes:   GetDiskStat(common.GetExecutionDirForTaskId(parser.AgentConfig.AgentDirectory, taskId)).Used,
	}
}
