package tui

import "github.com/charmbracelet/lipgloss"

const (
	pastelGreen = lipgloss.Color("#B2FBA5")
	pastelRed   = lipgloss.Color("#FF6961")
	pastelGrey  = lipgloss.Color("#E7E7E7")
	pastelBlue  = lipgloss.Color("#AEC6CF")
	darkGray    = lipgloss.Color("#767676")
)

var (
	inputStyle      = lipgloss.NewStyle().Foreground(pastelGreen)
	helperTextStyle = lipgloss.NewStyle().Foreground(darkGray)
	xStyle          = lipgloss.NewStyle().Foreground(pastelRed)
	taskText        = lipgloss.NewStyle().Foreground(pastelGreen)
	cursorStyle     = lipgloss.NewStyle().Foreground(pastelGreen)
)
