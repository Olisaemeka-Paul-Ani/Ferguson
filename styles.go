package main

import "github.com/charmbracelet/lipgloss"

const (
	colorAccent  = lipgloss.Color("#B45309")
	colorText    = lipgloss.Color("#D0D0D0")
	colorBorder  = lipgloss.Color("#444444")
	colorFooter  = lipgloss.Color("#666666")
	colorHard    = lipgloss.Color("#FF4444")
	colorNeutral = lipgloss.Color("#FFD700")
	colorEasy    = lipgloss.Color("#00FF88")
)

var (
	baseStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Padding(0, 1)

	paneStyle = baseStyle.
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder)

	activePaneStyle = paneStyle.
				BorderForeground(colorAccent)

	titleStyle = baseStyle.
			Bold(true).
			Foreground(colorAccent)

	footerStyle = baseStyle.
			Foreground(colorFooter)

	fdrHardStyle    = baseStyle.Foreground(colorHard)
	fdrNeutralStyle = baseStyle.Foreground(colorNeutral)
	fdrEasyStyle    = baseStyle.Foreground(colorEasy)
)
