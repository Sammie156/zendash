package widgets

import "github.com/charmbracelet/lipgloss"

type styles struct {
	Title       lipgloss.Style
	DetailKey   lipgloss.Style
	DetailValue lipgloss.Style
	Subheading  lipgloss.Style
	Container   lipgloss.Style
}

func newStyles() styles {
	return styles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("202")),

		DetailKey: lipgloss.NewStyle().
			Bold(true).
			Width(10),

		DetailValue: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")),

		Container: lipgloss.NewStyle().
			Border(lipgloss.ASCIIBorder()).
			Padding(1, 2),
	}
}
