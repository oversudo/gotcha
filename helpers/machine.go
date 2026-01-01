package helpers

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetGPUInfo() string {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("lspci")
	case "darwin":
		cmd = exec.Command("system_profiler", "SPDisplaysDataType")
	case "windows":
		cmd = exec.Command("powershell", "-Command", "Get-CimInstance -ClassName Win32_VideoController | Select-Object -ExpandProperty Name")
	default:
		return ""
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return parseGPUName(string(output))
}

func parseGPUName(output string) string {
	lines := strings.Split(output, "\n")

	switch runtime.GOOS {
	case "linux":
		// Look for VGA or 3D controller
		for _, line := range lines {
			if strings.Contains(line, "VGA") || strings.Contains(line, "3D controller") {
				parts := strings.Split(line, ": ")
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1])
				}
			}
		}
	case "darwin":
		// Look for Chipset Model
		for _, line := range lines {
			if strings.Contains(line, "Chipset Model:") {
				parts := strings.Split(line, ": ")
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1])
				}
			}
		}
	default:
		return strings.TrimSpace(output)
	}

	return "GPU not found"
}

func GetCPUInfo() string {
	info, err := cpu.Info()
	physicalCores, _ := cpu.Counts(false)
	logicalCores, _ := cpu.Counts(true)
	if err != nil {
		fmt.Printf("Error getting CPU info: %v\n", err)
		return ""
	}
	modelName := info[0].ModelName
	if modelName == "Undefined" {
		modelName = info[0].VendorID
	}

	if len(info) > 0 {
		return fmt.Sprintf("%s (%d Cores) (%d Threads)", modelName, physicalCores, logicalCores)
	}

	return ""
}
