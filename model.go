package main

import (
	"fmt"

	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width      int
	Height     int
	ActivePane int
	WillQuit   bool
	Squad      []fpl.Player
}

func (m Model) Init() tea.Cmd {
	return FetchPlayersCmd()

}

type TeamSheet struct {
	Players []fpl.Player
	Err     error
}

func FetchPlayersCmd() tea.Cmd {
	return func() tea.Msg {
		SheetData, err := fpl.FetchAllPlayers()
		if err != nil {
			return TeamSheet{Err: err}
		}
		return TeamSheet{Players: SheetData}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.WillQuit = true
			return m, tea.Quit

		}
	case TeamSheet:
		m.Squad = msg.Players

	}

	return m, nil

}

func (m Model) View() string {
	return fmt.Sprintf("Welcome to Ferguson. Squad loaded: %d players. Press q to quit", len(m.Squad))
}
