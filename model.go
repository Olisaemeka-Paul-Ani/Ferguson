package main

import (
	"strconv"
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
	Id           = "Identification"
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
		}).WithRows([]table.Row{}).WithPageSize(15).Focused(true),
		SparklineGraph: make(map[int][]fpl.Points),
	}
}

type Model struct {
	Width          int
	Height         int
	ActivePane     int
	WillQuit       bool
	ShowDetail     bool
	VerdictText    string
	Squad          []fpl.Player
	Fixtures       []fpl.Fixture
	RevealedChars  int
	LoadingIndex   int
	SquadErr       error
	FixtureErr     error
	VerdictErr     error
	FormErr        error
	simpleTable    table.Model
	SparklineGraph map[int][]fpl.Points
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(FetchPlayersCmd(), FetchFixturesCmd(), FetchVerdictCmd(), loadingTickCmd())

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

type FormSheet struct {
	id   int
	Err  error
	Form []fpl.Points
}

func FetchFormCmd(id int) tea.Cmd {
	return func() tea.Msg {
		FormData, err := fpl.FetchFormData(id)
		if err != nil {
			return FormSheet{Err: err}
		}
		result := FormSheet{id: id, Form: FormData}
		return result
	}
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

type loadingTickMsg time.Time

func loadingTickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*2000, func(t time.Time) tea.Msg {
		return loadingTickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.simpleTable, cmd = m.simpleTable.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.WillQuit = true
			return m, tea.Quit

		case "up", "down":
			if m.ShowDetail == true {
				id := m.simpleTable.HighlightedRow().Data[Id].(int)
				if m.SparklineGraph[id] != nil {
					return m, nil
				}
				return m, FetchFormCmd(id)
			}
			return m, nil

		case "enter":
			m.ShowDetail = true
			id := m.simpleTable.HighlightedRow().Data[Id].(int)
			if m.SparklineGraph[id] != nil {
				return m, nil
			}
			return m, FetchFormCmd(id)

		case "esc":
			m.ShowDetail = false

		case "right":
			m.ActivePane = (m.ActivePane + 1) % 3

		case "left":
			m.ActivePane = (m.ActivePane - 1 + 3) % 3

		}

	case FormSheet:
		if msg.Err != nil {
			m.FormErr = msg.Err
		} else {
			m.SparklineGraph[msg.id] = msg.Form
		}

	case TeamSheet:

		if msg.Err != nil {
			m.SquadErr = msg.Err
		} else {
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
					Id:           msg.Players[i].Identification,
				}
				rows = append(rows, table.NewRow(rowData))
				i = i + 1
			}
			m.simpleTable = m.simpleTable.WithRows(rows)
			m.Squad = msg.Players
		}

	case FixtureSheet:
		if msg.Err != nil {
			m.FixtureErr = msg.Err
		} else {
			m.Fixtures = msg.Fixtures
		}

	case VerdictSheet:
		if msg.Err != nil {
			m.VerdictErr = msg.Err
		} else {
			m.VerdictText = msg.Verdict
			return m, tickCmd()
		}

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

	case loadingTickMsg:
		m.LoadingIndex = (m.LoadingIndex + 1) % len(ui.LoadingQuotes)
		return m, loadingTickCmd()

	}

	return m, tea.Batch(cmds...)

}

func (m Model) View() string {

	gotSquad := len(m.Squad) > 0
	gotFixtures := len(m.Fixtures) > 0
	gotVerdict := len(m.VerdictText) > 0
	var GroupFirstFive map[int][]fpl.Fixture
	var FormatFixtures string
	var squadPane string
	var fixturesPane string
	var VerdictView string
	var HighLightedPane string

	if !gotSquad {
		if m.SquadErr != nil {
			return paneStyle.Render("Error: " + m.SquadErr.Error())
		}
		return paneStyle.Render("Loading squad...")
	} else if gotSquad {
		if m.ActivePane == 0 {
			squadPane = activePaneStyle.Render(m.simpleTable.View())
		} else {
			squadPane = paneStyle.Render(m.simpleTable.View())
		}
	}

	if !gotFixtures {
		if m.FixtureErr != nil {
			return paneStyle.Render("Error: " + m.FixtureErr.Error())
		}
		return paneStyle.Render("Loading fixtures...")
	} else if gotFixtures {
		HighlightedRow := m.simpleTable.HighlightedRow()
		HighlightedPlayer, ok := HighlightedRow.Data[PlayerName].(string)
		output := ""
		i := 0
		if !ok {
			output = "Error in asserting player Data"
			i = len(m.Squad)
		} else {
			i = 0
		}

		for i < len(m.Squad) {
			LookUp := ui.GetFixtureDict()
			PositionDict := ui.GetPlayerDict()
			current := m.Squad[i]
			if current.WebName == HighlightedPlayer {
				output = output + "\n"
				output = output + "Name: " + current.WebName + "\n"
				output = output + "Position: " + PositionDict[current.Position] + "\n"
				output = output + "Club: " + LookUp[current.Club] + "\n"
				output = output + "Cost: " + strconv.FormatFloat(float64(current.Cost)/10, 'f', 1, 64) + "\n"
				output = output + "TotalPoints " + strconv.Itoa(current.TotalPoints) + "\n"
				output = output + "GWPoints " + strconv.Itoa(current.GameweekPoints) + "\n"
				output = output + "Fitness: " + ui.StatusMap[current.Status] + " " + GetColor(current.ChanceOfPlaying) + " " + current.News + "\n"
				if m.FormErr != nil {
					output = output + "Form: Unable to Fetch Form Data." + "\n"
				} else {
					output = output + "Form: "
					if m.SparklineGraph[current.Identification] == nil {
						output = output + "Data Unavailable" + "\n"

					} else {
						output = output + ui.GetForm(m.SparklineGraph[current.Identification], FormMap) + "\n"
					}
				}
				output = output + "Fixtures: " + "\n"

				PlaceHolder := ui.FirstFive(ui.FixturesForClub(current.Club, ui.FindUpcomingMatches(m.Fixtures)))
				j := 0

				for j < len(PlaceHolder) {
					if current.Club == PlaceHolder[j].TeamHome {
						output = output + " " + "vs " + LookUp[PlaceHolder[j].TeamAway] + "(H) -" + "Difficulty " + fdrColorMap[PlaceHolder[j].TeamHomeDifficulty] + "\n"
					} else {
						output = output + " " + "vs " + LookUp[PlaceHolder[j].TeamHome] + "(A) -" + "Difficulty " + fdrColorMap[PlaceHolder[j].TeamAwayDifficulty] + "\n"
					}
					j = j + 1
				}

			}
			i = i + 1
		}

		GroupFirstFive = ui.GroupFirstFive(m.Fixtures)
		FormatFixtures = ui.FormatFixtures(GroupFirstFive, fdrColorMap)
		if m.ActivePane == 1 {
			fixturesPane = activePaneStyle.Render(FormatFixtures)
			HighLightedPane = ActiveHighlightStyle.Render(output)
		} else {
			fixturesPane = paneStyle.Render(FormatFixtures)
			HighLightedPane = HighlightStyle.Render(output)
		}

	}

	if !gotVerdict {
		if m.VerdictErr != nil {
			return paneStyle.Render("Error: " + m.VerdictErr.Error())
		}
		return paneStyle.Render("Loading Verdict...")
	} else if gotVerdict {
		if m.ActivePane == 2 {
			VerdictView = VerdictActivePaneStyle.Render(m.VerdictText[:m.RevealedChars])
		} else {
			VerdictView = verdictStyle.Render(m.VerdictText[:m.RevealedChars])
		}
	}

	if gotFixtures && gotSquad && gotVerdict {
		if m.ShowDetail == true {
			combined := lipgloss.JoinHorizontal(lipgloss.Top, squadPane, HighLightedPane, VerdictView)
			return combined
		} else {
			combined := lipgloss.JoinHorizontal(lipgloss.Top, squadPane, fixturesPane, VerdictView)
			return combined
		}

	}

	return ""

}
