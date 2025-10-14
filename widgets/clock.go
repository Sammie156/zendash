package widgets

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss" // <-- IMPORT lipgloss
)

// --- STYLES ---
// Define a new style for our clock's container.
var clockStyle = lipgloss.NewStyle().
	SetString("Clock\n"). // A title for the box
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")). // A nice purple
	Padding(1, 2)

// --- ASCII ART FONT ---
// A map to store our simple 3-line ASCII font for digits 0-9 and the colon.
var digitalFont = map[rune][]string{
	'0': {" 000 ", "0   0", " 000 "},
	'1': {"  1  ", "  1  ", "  1  "},
	'2': {" 222 ", "  2  ", " 222 "},
	'3': {" 333 ", "  33 ", " 333 "},
	'4': {"4  4 ", " 444 ", "   4 "},
	'5': {" 555 ", " 55  ", " 555 "},
	'6': {" 666 ", "666  ", " 666 "},
	'7': {" 777 ", "   7 ", "   7 "},
	'8': {" 888 ", " 888 ", " 888 "},
	'9': {" 999 ", "  999", " 999 "},
	':': {"     ", "  :  ", "     "},
}

// (tickMsg and tick() function remain the same)
type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// (ClockModel and NewClock remain the same)
type ClockModel struct {
	time time.Time
}

func NewClock() ClockModel {
	return ClockModel{time: time.Now()}
}

// (Init and Update methods remain the same)
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

// View is our NEW and IMPROVED render function.
func (m ClockModel) View() string {
	// Get the time in HH:MM:SS format
	timeStr := m.time.Format("15:04:05")

	// Create three string builders, one for each line of our ASCII art.
	var line1, line2, line3 strings.Builder

	// Iterate over each character in the time string (e.g., '1', '4', ':', '0', '5')
	for _, char := range timeStr {
		// Look up the ASCII art for the character
		font, ok := digitalFont[char]
		if ok {
			// Append the corresponding line of the font to our builders
			line1.WriteString(font[0])
			line2.WriteString(font[1])
			line3.WriteString(font[2])
		}
	}

	// Join the three lines of ASCII art together with newlines
	asciiClock := fmt.Sprintf("%s\n%s\n%s", line1.String(), line2.String(), line3.String())

	// Use our lipgloss style to render the final output with a border!
	return clockStyle.Render(asciiClock)
}
