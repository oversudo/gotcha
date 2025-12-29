package helpers

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetCPUInfo() string {
	info, err := cpu.Info()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v\n", err)
		return ""
	}

	if len(info) > 0 {
		return fmt.Sprintf("%s (%d Cores)", info[0].ModelName, info[0].Cores)
	}

	return ""
}
