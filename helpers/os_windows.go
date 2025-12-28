package helpers

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func GetOSInfo() string {
	osvi := windows.RtlGetVersion()
	var windowsVersion string
	switch {
	case osvi.MajorVersion == 10 && osvi.BuildNumber >= 22000:
		windowsVersion = "Windows 11"
	case osvi.MajorVersion == 10:
		windowsVersion = "Windows 10"
	case osvi.MajorVersion == 6 && osvi.MinorVersion == 3:
		windowsVersion = "Windows 8.1"
	case osvi.MajorVersion == 6 && osvi.MinorVersion == 2:
		windowsVersion = "Windows 8"
	case osvi.MajorVersion == 6 && osvi.MinorVersion == 1:
		windowsVersion = "Windows 7"
	default:
		windowsVersion = fmt.Sprintf("Windows NT %d.%d", osvi.MajorVersion, osvi.MinorVersion)
	}

	return fmt.Sprintf("%s (%d)", windowsVersion, osvi.BuildNumber)
}

func GetKernelVersion() string {
	return "11"
}
