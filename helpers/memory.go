package helpers

import (
	"github.com/shirou/gopsutil/v4/mem"
)

func GetMemory() mem.VirtualMemoryStat {
	v, err := mem.VirtualMemory()
	if err != nil {
		return mem.VirtualMemoryStat{}
	}

	return *v
}
