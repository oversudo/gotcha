package helpers

import (
	"os/exec"
	"strings"
)

func NumberOfPackages() map[string]int {
	result := make(map[string]int)
	var cmd *exec.Cmd

	if commandExists("dpkg") {
		cmd = exec.Command("sh", "-c", "dpkg -l | grep '^ii'")
		result["dpkg"] = linesToCount(cmd)
	}
	if commandExists("rpm") {
		cmd = exec.Command("rpm", "-qa")
		result["rpm"] = linesToCount(cmd)
	}
	if commandExists("pacman") {
		cmd = exec.Command("pacman", "-Q")
		result["pacman"] = linesToCount(cmd)
	}
	if commandExists("brew") {
		cmd = exec.Command("brew", "list", "--formula")
		result["brew"] = linesToCount(cmd)
		cmd = exec.Command("brew", "list", "--cask")
		result["brew-cask"] = linesToCount(cmd)
	}

	return result
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func linesToCount(cmd *exec.Cmd) int {
	output, err := cmd.Output()
	if err != nil {
		return 0
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return len(lines)
}
