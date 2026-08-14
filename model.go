package main

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Width      int
	Height     int
	ActivePane int
	WillQuit   bool
}

func (m Model) Init() tea.Cmd {
	return nil
}
