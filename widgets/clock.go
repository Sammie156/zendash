package widgets

// TODO: Recreate the Clock Widget using lipgloss
import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type ClockModel struct {
	time time.Time
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
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
	timeStr := m.time.Format("15:01:05")

	return timeStr
}
