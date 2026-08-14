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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.WillQuit = true
			return m, tea.Quit

		}

	}

	return m, nil

}

func (m Model) View() string {
	return "Welcome to Ferguson. Press q to quit"
}
