package main

import (
	"fmt"
	"os"
	"zendash/widgets"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type mainModel struct {
	width  int
	height int
	clock  widgets.ClockModel
}

func newMainModel(t_width int, t_height int) mainModel {
	return mainModel{
		width:  t_width,
		height: t_height,
		clock:  widgets.NewClock(t_width, t_height),
	}
}

func (m mainModel) Init() tea.Cmd {
	return m.clock.Init()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			fmt.Println() // This is here so that after quiting, the last line doesn't get cut
			return m, tea.Quit
		}
	}

	var updatedClock widgets.ClockModel
	updatedClock, cmd = m.clock.Update(msg)
	m.clock = updatedClock

	return m, cmd
}

func (m mainModel) View() string {
	clockView := m.clock.View()

	return clockView
}

// TODO: Reformat the code and make it look better
// TODO: Call the clock widget and append it to the dashboard
func main() {
	fd := int(os.Stdout.Fd())

	// Check if standard input is a terminal
	if !term.IsTerminal(fd) {
		fmt.Println("Standard output is not a terminal. Zendash is made for for the CLI")
		return
	}

	width, height, err := term.GetSize(fd)

	if err != nil {
		fmt.Printf("Error getting terminal size %s\n", err)
		return
	}

	p := tea.NewProgram(newMainModel(width, height))

	var style = lipgloss.NewStyle().
		Padding(1).
		MarginLeft(width/2-15).
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.AdaptiveColor{Light: "201", Dark: "#75D"}).
		Border(lipgloss.ASCIIBorder(), true, true).
		Width(30).
		Align(lipgloss.Center)

	fmt.Println(style.Render("Welcome to ZenDash"))

	if _, err := p.Run(); err != nil {
		fmt.Printf("There was an error somewhere! %v", err)
		os.Exit(1)
	}
}
