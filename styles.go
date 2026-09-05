package main

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	colorAccent   = lipgloss.Color("#B45309")
	colorText     = lipgloss.Color("#D0D0D0")
	colorBorder   = lipgloss.Color("#444444")
	colorFooter   = lipgloss.Color("#666666")
	colorVeryHard = lipgloss.Color("#8B0000")
	colorHard     = lipgloss.Color("#FF4444")
	colorNeutral  = lipgloss.Color("#FFD700")
	colorEasy     = lipgloss.Color("#00FF88")
	colorVeryEasy = lipgloss.Color("#3399FF")
)

var fdrColorMap = map[int]string{
	1: fdrVeryEasyStyle.Render("■■■■■"),
	2: fdrEasyStyle.Render("■■■■■"),
	3: fdrNeutralStyle.Render("■■■■■"),
	4: fdrHardStyle.Render("■■■■■"),
	5: fdrVeryHardStyle.Render("■■■■■"),
}

var FtColorMap = map[int]string{
	0:   fdrVeryHardStyle.Render("□□□□□"),
	25:  fdrHardStyle.Render("■□□□□"),
	50:  fdrNeutralStyle.Render("■■■□□"),
	75:  fdrEasyStyle.Render("■■■■□"),
	100: fdrVeryEasyStyle.Render("■■■■■"),
}

func GetColor(arg *int) string {
	if arg == nil {
		return FtColorMap[100]
	} else {
		return FtColorMap[*arg]
	}
}

var (
	baseStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Padding(0, 1)

	paneStyle = baseStyle.
			Border(lipgloss.ThickBorder()).
			BorderForeground(colorBorder)

	activePaneStyle = paneStyle.
			BorderForeground(colorAccent)

	titleStyle = baseStyle.
			Bold(true).
			Foreground(colorAccent)

	footerStyle = baseStyle.
			Foreground(colorFooter)

	verdictStyle = paneStyle.
			Copy().
			Foreground(colorAccent).
			Width(40).
			MaxHeight(15)

	fdrVeryHardStyle = baseStyle.Foreground(colorVeryHard)
	fdrHardStyle     = baseStyle.Foreground(colorHard)
	fdrNeutralStyle  = baseStyle.Foreground(colorNeutral)
	fdrEasyStyle     = baseStyle.Foreground(colorEasy)
	fdrVeryEasyStyle = baseStyle.Foreground(colorVeryEasy)
)
