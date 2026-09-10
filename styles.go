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

var FormMap = map[int]string{
	-1: sparkBadStyle.Render("▁"),
	0:  sparkGoodStyle.Render("▁"),
	1:  sparkGoodStyle.Render("▁"),
	2:  sparkGoodStyle.Render("▂"),
	3:  sparkGoodStyle.Render("▂"),
	4:  sparkGoodStyle.Render("▃"),
	5:  sparkGoodStyle.Render("▃"),
	6:  sparkGoodStyle.Render("▅"),
	7:  sparkGoodStyle.Render("▅"),
	8:  sparkGoodStyle.Render("▅"),
	9:  sparkGoodStyle.Render("▇"),
	10: sparkGoodStyle.Render("▇"),
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

	loadingStyle = paneStyle.
			Width(80).
			Align(lipgloss.Center)

	activePaneStyle = paneStyle.
			BorderForeground(colorAccent)

	titleStyle = baseStyle.
			Bold(true).
			Foreground(colorAccent)

	HeadStyle = titleStyle.
			Copy().
			Align(lipgloss.Center).
			Width(80)

	footerStyle = baseStyle.
			Foreground(colorFooter)

	verdictStyle = paneStyle.
			Copy().
			Foreground(colorAccent).
			Width(40).
			MaxHeight(15)

	HighlightStyle = paneStyle.
			Copy().
			Width(40)

	ActiveHighlightStyle = paneStyle.
				Copy().
				BorderForeground(colorAccent).
				Width(40)

	VerdictActivePaneStyle = paneStyle.
				Copy().
				Foreground(colorAccent).
				BorderForeground(colorAccent).
				Width(40).
				MaxHeight(15)

	FixtureViewStyle = paneStyle.
				Copy().
				MaxHeight(15)

	ActiveFixtureViewStyle = activePaneStyle.
				Copy().
				MaxHeight(15)

	fdrVeryHardStyle = baseStyle.Foreground(colorVeryHard)
	fdrHardStyle     = baseStyle.Foreground(colorHard)
	fdrNeutralStyle  = baseStyle.Foreground(colorNeutral)
	fdrEasyStyle     = baseStyle.Foreground(colorEasy)
	fdrVeryEasyStyle = baseStyle.Foreground(colorVeryEasy)
	sparkBadStyle    = baseStyle.Foreground(colorHard)
	sparkGoodStyle   = baseStyle.Foreground(colorAccent)
)
