package renderer

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/oversudo/gofetch/helpers"
	"github.com/oversudo/gofetch/logo"
)

var (

	// Layout
	leftStyle = lipgloss.NewStyle().
		Width(50).
		Padding(1, 2)

	rightStyle = lipgloss.NewStyle().
		Padding(1, 2)

	// Color scheme
	accentColor = lipgloss.Color("#00D9FF")
	labelColor  = lipgloss.Color("#7AA2F7")
	textColor   = lipgloss.Color("#C0CAF5")
	mutedColor  = lipgloss.Color("#565F89")

	// Styles
	titleStyle = lipgloss.NewStyle().
		Foreground(accentColor).
		Bold(true)

	labelStyle = lipgloss.NewStyle().
		Foreground(labelColor).
		Bold(true)

	valueStyle = lipgloss.NewStyle().
		Foreground(textColor)

	separatorStyle = lipgloss.NewStyle().
		Foreground(mutedColor)
)

func Render() {
	user := helpers.GetUsername()
	host := helpers.GetHostname()
	userHost := titleStyle.Render(fmt.Sprintf("%s@%s", user, host))
	separator := separatorStyle.Render(strings.Repeat("─", len(user)+len(host)+1))
	infoLines := []string{
		userHost,
		separator,
	}
	infoLines = append(infoLines, renderInfoLine("OS", helpers.GetOSInfo()))
	infoLines = append(infoLines, renderInfoLine("Uptime", helpers.GetUptime()))
	if runtime.GOOS != "windows" {
		infoLines = append(infoLines, renderInfoLine("Kernel", helpers.GetKernelVersion()))
		var packagesString []string
		for pkg, count := range helpers.NumberOfPackages() {
			packagesString = append(packagesString, fmt.Sprintf("%d (%s)", count, pkg))
		}
		infoLines = append(infoLines, renderInfoLine("Packages", strings.Join(packagesString, ", ")))
	}
	infoLines = append(infoLines, renderInfoLine("Shell", helpers.GetShellInfo()))
	infoLines = append(infoLines, renderInfoLine("Public IP", helpers.GetExternalIP()))
	infoLines = append(infoLines, renderInfoLine("Private IPs", strings.Join(helpers.GetLocalIPs(), ", ")))
	infoLines = append(infoLines, renderInfoLine("CPU", helpers.GetCPUInfo()))
	infoLines = append(infoLines, renderInfoLine("GPU", helpers.GetGPUInfo()))
	memory := helpers.GetMemory()
	infoLines = append(infoLines, renderInfoLine("Memory", fmt.Sprintf("%d / %d MB (%.2f%%)",
		memory.Used/1024/1024, memory.Total/1024/1024, memory.UsedPercent)))
	var resolutions []string
	for _, display := range helpers.GetDisplays() {
		if display.Primary {
			resolutions = append(resolutions, fmt.Sprintf("%s (Primary)", display.Resolution))
		} else {
			resolutions = append(resolutions, fmt.Sprintf("%s", display.Resolution))
		}
	}
	infoLines = append(infoLines, renderInfoLine("Resolution", strings.Join(resolutions, ", ")))
	leftContent := logo.DEFAULT
	rightContent := strings.Join(infoLines, "\n")

	left := leftStyle.Render(leftContent)
	right := rightStyle.Render(rightContent)

	ui := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	fmt.Println(ui)
}

func renderInfoLine(label, value string) string {
	sep := separatorStyle.Render(": ")
	return labelStyle.Render(label) + sep + valueStyle.Render(value)
}
