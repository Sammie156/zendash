package main

import (
	"fmt"
	"os"
	"zendash/widgets"

	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	clock widgets.ClockModel
}

func initialModel() mainModel {
	return mainModel{
		clock: widgets.NewClock(),
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
			return m, tea.Quit
		}
	}

	var updatedClock widgets.ClockModel
	updatedClock, cmd = m.clock.Update(msg)
	m.clock = updatedClock

	return m, cmd
}

func (m mainModel) View() string {
	s := "Welcome to ZenDash \n\n"

	s += m.clock.View() + "\n\n"

	s += "Press q to exit. \n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error : %v", err)
		os.Exit(1)
	}
}
