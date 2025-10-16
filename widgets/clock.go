package widgets

// TODO: Recreate the Clock Widget using lipgloss
import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type ClockModel struct {
	time   time.Time
	zone   string
	styles styles
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewClock(t_width int, t_height int) ClockModel {
	zoneStr := time.Now().Local().String()[28:]

	return ClockModel{time: time.Now(), zone: zoneStr, styles: newStyles("CLOCK")}
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
	title := m.styles.Title.Render("")

	time_key := m.styles.DetailKey.Render("Current Time:")
	time_value := m.styles.DetailValue.Render(m.time.Format("15:04:05"))
	time_line := fmt.Sprintf("%s %s", time_key, time_value)

	zone_key := m.styles.DetailKey.Render("Time Zone:")
	zone_value := m.styles.DetailValue.Render(m.zone)
	zone_line := fmt.Sprintf("%s %s", zone_key, zone_value)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		" ",
		time_line,
		zone_line,
		" ",
	)

	return m.styles.Container.Render(content)
}
