package helpers

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"golang.org/x/sys/unix"
)

func GetOSInfo() string {
	osName := getOSName()
	arch := getArchName()
	return fmt.Sprintf("%s %s", osName, arch)
}

func GetKernelVersion() string {
	var utsname unix.Utsname
	if err := unix.Uname(&utsname); err != nil {
		return ""
	}
	release := string(utsname.Release[:])
	return release
}

func getOSName() string {
	switch runtime.GOOS {
	case "linux":
		return getLinuxOrBSDName()
	case "darwin":
		return getDarwinName()
	case "windows":
		return getWindowsName()
	case "freebse", "openbsd", "netbsd":
		return getLinuxOrBSDName()
	default:
		return runtime.GOOS
	}
}

func getLinuxOrBSDName() string {
	if out, err := exec.Command("uname", "-sr").Output(); err == nil {
		return strings.TrimSpace(string(out))
	}
	return "Linux"
}

func getDarwinName() string {
	productName, _ := exec.Command("sw_vers", "-productName").Output()
	version, _ := exec.Command("sw_vers", "-productVersion").Output()

	return fmt.Sprintf("%s %s", strings.TrimSpace(string(productName)), strings.TrimSpace(string(version)))
}

func getWindowsName() string {
	// Use wmic to get OS name
	out, err := exec.Command("wmic", "os", "get", "Caption", "/value").Output()
	if err == nil {
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "Caption=") {
				caption := strings.TrimPrefix(line, "Caption=")
				return strings.TrimSpace(caption)
			}
		}
	}

	// Fallback to ver command
	out, err = exec.Command("cmd", "/c", "ver").Output()
	if err == nil {
		return strings.TrimSpace(string(out))
	}

	return "Windows"
}

func getArchName() string {
	archMap := map[string]string{
		"amd64":   "x86_64",
		"386":     "i686",
		"arm64":   "aarch64",
		"arm":     "armv7l",
		"ppc64le": "ppc64le",
		"s390x":   "s390x",
		"mips64":  "mips64",
	}

	if arch, ok := archMap[runtime.GOARCH]; ok {
		return arch
	}
	return runtime.GOARCH
}
