package main

import (
	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
	"github.com/Olisaemeka-Paul-Ani/ferguson/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width      int
	Height     int
	ActivePane int
	WillQuit   bool
	Squad      []fpl.Player
	Fixtures   []fpl.Fixture
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(FetchPlayersCmd(), FetchFixturesCmd())

}

type TeamSheet struct {
	Players []fpl.Player
	Err     error
}

type FixtureSheet struct {
	Fixtures []fpl.Fixture
	Err      error
}

func FetchFixturesCmd() tea.Cmd {
	return func() tea.Msg {
		SheetData, err := fpl.FetchAllFixtures()
		if err != nil {
			return FixtureSheet{Err: err}
		}
		result := FixtureSheet{Fixtures: SheetData}
		return result
	}
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

	case FixtureSheet:
		m.Fixtures = msg.Fixtures
	}

	return m, nil

}

func (m Model) View() string {
	var CleanedData [][]string
	var FormattedData string
	gotSquad := len(m.Squad) > 0
	gotFixtures := len(m.Fixtures) > 0
	var GroupFirstFive map[int][]fpl.Fixture
	var FormatFixtures string
	var squadPane string
	var fixturesPane string

	if !gotSquad {
		return paneStyle.Render("Loading squad...")
	} else if gotSquad {

		CleanedData = ui.CleanData(m.Squad)
		FormattedData = ui.FormatData(CleanedData)
		squadPane = activePaneStyle.Render(FormattedData)

	}

	if !gotFixtures {
		return paneStyle.Render("Loading fixtures...")
	} else if gotFixtures {
		GroupFirstFive = ui.GroupFirstFive(m.Fixtures)
		FormatFixtures = ui.FormatFixtures(GroupFirstFive, fdrColorMap)
		fixturesPane = paneStyle.Render(FormatFixtures)

	}
	if gotFixtures && gotSquad {
		combined := lipgloss.JoinHorizontal(lipgloss.Top, squadPane, fixturesPane)
		return combined
	}

	return ""

}
