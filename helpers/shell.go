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

	if out, err := exec.Command(shell, "--version").Output(); err == nil {
		return strings.TrimSpace(string(out))
	}
	return ""
}
