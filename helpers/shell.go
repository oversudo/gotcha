package helpers

import (
	"fmt"
	"os"
	"strings"
)

func GetShellInfo() string {
	return fmt.Sprintf("%s %s", getShell(), getShellVersion())
}

func getShell() string {
	shellPath := os.Getenv("SHELL")
	lastIndex := strings.LastIndex(shellPath, "/")
	if lastIndex != -1 {
		return shellPath[lastIndex+1:]
	}
	return ""
}

func getShellVersion() string {
	return "1"
}