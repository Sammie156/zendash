package widgets

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type ClockModel struct {
	time time.Time
}

func NewClock() ClockModel {
	return ClockModel{time: time.Now()}
}

func (m ClockModel) Init() tea.Cmd {
	return tick()
}

func (m ClockModel) Update(msg tea.Msg) (ClockModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.time = time.Time(msg)
		return m, tick()
	}

	return m, nil
}

func (m ClockModel) View() string {
	return fmt.Sprintf("It is currently: %s", m.time.Format("15:04:05 PM"))
}
