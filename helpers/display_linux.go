package helpers

import (
	"os/exec"
	"regexp"
	"strings"
)

type Display struct {
	Resolution  string
	RefreshRate float64
	Primary     bool
}

func GetDisplays() []Display {
	return getXrandrDisplays()
}

func getXrandrDisplays() []Display {
	output, err := exec.Command("xrandr", "--query").Output()
	if err != nil {
		return nil
	}

	var displays []Display
	lines := strings.Split(string(output), "\n")
	pattern := `\d+x\d+`
	re := regexp.MustCompile(pattern)
	for _, line := range lines {
		if strings.Contains(line, " connected") {
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}

			display := Display{
				Primary:    strings.Contains(line, "primary"),
				Resolution: re.FindString(line),
			}

			displays = append(displays, display)
		}

	}

	return displays

}
