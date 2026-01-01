package renderer

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/oversudo/gotcha/helpers"
	"github.com/oversudo/gotcha/logo"
)

type Line struct {
	Key   string
	Value string
}

func Render() {
	fieldOrder := []string{
		"OS",
		"Uptime",
		"Kernel",   // Skip on Windows
		"Packages", // Skip on Windows
		"Shell",
		"Public IP",
		"Private IPs",
		"CPU",
		"GPU",
		"Memory",
		"Resolution",
	}

	user := helpers.GetUsername()
	host := helpers.GetHostname()
	userHost := titleStyle.Render(fmt.Sprintf("%s@%s", user, host))
	separator := separatorStyle.Render(strings.Repeat("─", len(user)+len(host)+1))
	infoLines := []string{
		userHost,
		separator,
	}

	outputCh := make(chan Line)
	var wg sync.WaitGroup

	go func() {
		wg.Wait() // Blocks until wg counter = 0
		close(outputCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "OS", Value: helpers.GetOSInfo()}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "Uptime", Value: helpers.GetUptime()}
	}()

	if runtime.GOOS != "windows" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			outputCh <- Line{Key: "Kernel", Value: helpers.GetKernelVersion()}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			var packagesString []string
			for pkg, count := range helpers.NumberOfPackages() {
				packagesString = append(packagesString, fmt.Sprintf("%d (%s)", count, pkg))
			}
			outputCh <- Line{Key: "Packages", Value: strings.Join(packagesString, ", ")}
		}()

	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "Shell", Value: helpers.GetShellInfo()}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "Public IP", Value: helpers.GetExternalIP()}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "Private IPs", Value: strings.Join(helpers.GetLocalIPs(), ", ")}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "CPU", Value: helpers.GetCPUInfo()}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outputCh <- Line{Key: "GPU", Value: helpers.GetGPUInfo()}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		memory := helpers.GetMemory()
		outputCh <- Line{Key: "Memory", Value: fmt.Sprintf("%d / %d MB (%.2f%%)",
			memory.Used/1024/1024, memory.Total/1024/1024, memory.UsedPercent)}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		var resolutions []string
		for _, display := range helpers.GetDisplays() {
			if display.Primary {
				resolutions = append(resolutions, fmt.Sprintf("%s (Primary)", display.Resolution))
			} else {
				resolutions = append(resolutions, fmt.Sprintf("%s", display.Resolution))
			}
		}
		outputCh <- Line{Key: "Resolution", Value: strings.Join(resolutions, ", ")}
	}()

	results := make(map[string]string)
	for line := range outputCh {
		results[line.Key] = line.Value
	}

	for _, field := range fieldOrder {
		if value, ok := results[field]; ok {
			infoLines = append(infoLines, renderInfoLine(field, value))
		}
	}

	leftContent := logo.GetLogo()
	rightContent := strings.Join(infoLines, "\n")

	left := leftStyle.Render(leftContent)
	right := rightStyle.Render(rightContent)

	ui := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	fmt.Println(ui)
}
