package main

import (
	"time"

	"github.com/Olisaemeka-Paul-Ani/ferguson/ai"
	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"

	"github.com/Olisaemeka-Paul-Ani/ferguson/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const (
	PlayerName   = "Player"
	PositionName = "Pos"
	ClubName     = "Club"
	Price        = "Cost"
	TotalPoints  = "TotalPoints"
	GWPoints     = "GW Points"
)

func NewModel() Model {
	return Model{
		simpleTable: table.New([]table.Column{
			table.NewColumn(PlayerName, "Name", 16),
			table.NewColumn(PositionName, "Pos", 5),
			table.NewColumn(ClubName, "Club", 6),
			table.NewColumn(Price, "cost", 6),
			table.NewColumn(TotalPoints, "TotalPoints", 12),
			table.NewColumn(GWPoints, "GW Points", 10),
		}).WithRows([]table.Row{}),
	}
}

type Model struct {
	Width         int
	Height        int
	ActivePane    int
	WillQuit      bool
	VerdictText   string
	Squad         []fpl.Player
	Fixtures      []fpl.Fixture
	RevealedChars int
	SquadErr      error
	FixtureErr    error
	VerdictErr    error
	simpleTable   table.Model
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(FetchPlayersCmd(), FetchFixturesCmd(), FetchVerdictCmd())

}

type TeamSheet struct {
	Players []fpl.Player
	Err     error
}

type FixtureSheet struct {
	Fixtures []fpl.Fixture
	Err      error
}

type VerdictSheet struct {
	Verdict string
	Err     error
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

func FetchVerdictCmd() tea.Cmd {
	return func() tea.Msg {
		SheetData, err := ai.FallBackFunction()
		if err != nil {
			return VerdictSheet{Err: err}
		}
		result := VerdictSheet{Verdict: SheetData}
		return result
	}
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*25, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
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

		if msg.Err != nil {
			m.SquadErr = msg.Err
		}
		i := 0
		var rows []table.Row
		for i < len(msg.Players) {
			rowData := table.RowData{
				PlayerName:   msg.Players[i].WebName,
				PositionName: msg.Players[i].Position,
				ClubName:     msg.Players[i].Club,
				Price:        msg.Players[i].Cost,
				TotalPoints:  msg.Players[i].TotalPoints,
				GWPoints:     msg.Players[i].GameweekPoints,
			}
			rows = append(rows, table.NewRow(rowData))
			i = i + 1
		}
		m.simpleTable = m.simpleTable.WithRows(rows)
		m.Squad = msg.Players

	case FixtureSheet:
		if msg.Err != nil {
			m.FixtureErr = msg.Err
		}
		m.Fixtures = msg.Fixtures

	case VerdictSheet:
		if msg.Err != nil {
			m.VerdictErr = msg.Err
		}
		m.VerdictText = msg.Verdict
		return m, tickCmd()

	case tickMsg:
		var remaining int
		remaining = len(m.VerdictText) - m.RevealedChars
		if remaining < 3 {
			m.RevealedChars = m.RevealedChars + remaining
		} else {
			m.RevealedChars = m.RevealedChars + 3
		}

		if m.RevealedChars < len(m.VerdictText) {
			return m, tickCmd()
		}
	}

	return m, nil

}

func (m Model) View() string {
	var CleanedData [][]string
	var FormattedData string
	gotSquad := len(m.Squad) > 0
	gotFixtures := len(m.Fixtures) > 0
	gotVerdict := len(m.VerdictText) > 0
	var GroupFirstFive map[int][]fpl.Fixture
	var FormatFixtures string
	var squadPane string
	var fixturesPane string
	var VerdictView string

	if !gotSquad {
		if m.SquadErr != nil {
			return paneStyle.Render("Error: " + m.SquadErr.Error())
		}
		return paneStyle.Render("Loading squad...")
	} else if gotSquad {

		CleanedData = ui.CleanData(m.Squad)
		FormattedData = ui.FormatData(CleanedData)
		squadPane = activePaneStyle.Render(FormattedData)

	}

	if !gotFixtures {
		if m.FixtureErr != nil {
			return paneStyle.Render("Error: " + m.FixtureErr.Error())
		}
		return paneStyle.Render("Loading fixtures...")
	} else if gotFixtures {
		GroupFirstFive = ui.GroupFirstFive(m.Fixtures)
		FormatFixtures = ui.FormatFixtures(GroupFirstFive, fdrColorMap)
		fixturesPane = paneStyle.Render(FormatFixtures)

	}

	if !gotVerdict {
		if m.VerdictErr != nil {
			return paneStyle.Render("Error: " + m.VerdictErr.Error())
		}
		return paneStyle.Render("Loading Verdict...")
	} else if gotVerdict {
		VerdictView = verdictStyle.Render(m.VerdictText[:m.RevealedChars])
	}

	if gotFixtures && gotSquad && gotVerdict {
		combined := lipgloss.JoinHorizontal(lipgloss.Top, squadPane, fixturesPane, VerdictView)
		return combined
	}

	return ""

}
