package main

import (
	"fmt"
	"os"
	"zendash/widgets"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss" // <-- IMPORT lipgloss
)

// (mainModel and newMainModel remain the same)
type mainModel struct {
	clock widgets.ClockModel
}

func newMainModel() mainModel {
	return mainModel{
		clock: widgets.NewClock(),
	}
}

// (Init and Update remain the same)
func (m mainModel) Init() tea.Cmd {
	return m.clock.Init()
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var updatedClock widgets.ClockModel
	updatedClock, cmd = m.clock.Update(msg)
	m.clock = updatedClock

	return m, cmd
}

// View now uses lipgloss to join widgets together.
func (m mainModel) View() string {
	// Get the string output of our clock widget's View() method.
	clockView := m.clock.View()

	// An example of a placeholder for a future widget.
	// We'll replace this later with a real to-do list.
	placeholder := lipgloss.NewStyle().
		SetString("Future To-Do Widget").
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Render("...")

	// Use lipgloss.JoinVertical to stack our widgets.
	// The first argument is the alignment (0.5 means center).
	mainView := lipgloss.JoinVertical(
		lipgloss.Center,
		clockView,
		placeholder,
	)

	return mainView
}

func main() {
	p := tea.NewProgram(newMainModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
