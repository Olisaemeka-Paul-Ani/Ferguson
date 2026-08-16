package main

import (
	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
	"github.com/Olisaemeka-Paul-Ani/ferguson/ui"
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
		result := TeamSheet{Players: SheetData}
		return result
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
	CleanedData := ui.CleanData(m.Squad)
	FormattedData := ui.FormatData(CleanedData)
	return activePaneStyle.Render(FormattedData)
}
