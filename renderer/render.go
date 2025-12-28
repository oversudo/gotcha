package renderer

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/oversudo/gofetch/helpers"
)

var (
	// Color scheme
	accentColor    = lipgloss.Color("#00D9FF")
	secondaryColor = lipgloss.Color("#FF6AC1")
	labelColor     = lipgloss.Color("#7AA2F7")
	textColor      = lipgloss.Color("#C0CAF5")
	mutedColor     = lipgloss.Color("#565F89")

	// Styles
	titleStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true).
			MarginBottom(1)

	labelStyle = lipgloss.NewStyle().
			Foreground(labelColor).
			Bold(true)

	valueStyle = lipgloss.NewStyle().
			Foreground(textColor)

	separatorStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	asciiStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true)

	colorBarStyle = lipgloss.NewStyle().
			Bold(true)
)

// ASCII art for different OSes
var asciiArt = []string{
	"        ___       ",
	"       (.. |      ",
	"       (<> |      ",
	"      / __  \\     ",
	"     ( /  \\ /|    ",
	"    _/\\ __)/_)    ",
	"    \\/-____\\/     ",
}

func Render() {
	user := helpers.GetUsername()
	host := helpers.GetHostname()
	userHost := titleStyle.Render(fmt.Sprintf("%s@%s", user, host))
	separator := separatorStyle.Render(strings.Repeat("─", len(user)+len(host)+1))
	infoLines := []string{
		userHost,
		separator,
		renderInfoLine("OS: ", helpers.GetOSInfo()),
		renderInfoLine("Uptime: ", helpers.GetUptime()),
	}
	for _, value := range infoLines {
		fmt.Printf("%s\n", value)
	}
}

func renderInfoLine(label, value string) string {
	sep := separatorStyle.Render(" ")
	return labelStyle.Render(label) + sep + valueStyle.Render(value)
}
