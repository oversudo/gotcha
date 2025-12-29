//go:build !windows

package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetShellInfo() string {
	return fmt.Sprintf("%s", getShellVersion())
}

func getShell() (string, error) {
	shellPath := os.Getenv("SHELL")
	lastIndex := strings.LastIndex(shellPath, "/")
	if lastIndex != -1 {
		return shellPath[lastIndex+1:], nil
	}
	return "", fmt.Errorf("shell not found")
}

func getShellVersion() string {
	shell, err := getShell()
	if err != nil {
		return ""
	}

	switch shell {
		case "bash":
			cmd := exec.Command("bash", "-c", "bash --version | head -1 | cut -d ' ' -f 4-")
			output, err := cmd.Output()
			if err != nil {
				return ""
			}
			return fmt.Sprintf("bash %s",strings.TrimSpace(string(output)))
		default:
			if out, err := exec.Command(shell, "--version").Output(); err == nil {
				if strings.HasPrefix(string(out),shell){
					return strings.TrimSpace(string(out))
				} else {
					return fmt.Sprintf("%s %s", shell, strings.TrimSpace(string(out)))
				}
			}
	}

	return ""
}
