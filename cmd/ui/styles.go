package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	PrimaryColor   = lipgloss.Color("#7D5FFF") // Purple
	SecondaryColor = lipgloss.Color("#5EDEFF") // Cyan
	BorderColor    = lipgloss.Color("#4A4A4A") // Subtle gray
	TitleColor     = lipgloss.Color("#FFFFFF")
	ErrorColor     = lipgloss.Color("#FF6B6B")
	SuccessColor   = lipgloss.Color("#4CD964")
	WarningColor   = lipgloss.Color("#FFCC00")

	TitleStyle = lipgloss.NewStyle().
			Foreground(TitleColor).
			Bold(true).
			Padding(0, 1)

	SectionStyle = lipgloss.NewStyle().
			Foreground(PrimaryColor).
			MarginTop(1).
			MarginBottom(1).
			Bold(true)

	BorderBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor).
			Padding(1, 2).
			Margin(1, 0)

	MutedText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0"))
)