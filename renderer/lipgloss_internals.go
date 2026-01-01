package renderer

import "github.com/charmbracelet/lipgloss"

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

func renderInfoLine(label, value string) string {
	sep := separatorStyle.Render(": ")
	return labelStyle.Render(label) + sep + valueStyle.Render(value)
}
