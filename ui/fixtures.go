package ui

import (
	"strconv"

	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
)

func getFixtureDict() map[int]string {
	FixtureDict := map[int]string{
		1:  "Arsenal",
		2:  "Aston Villa",
		3:  "Bournemouth",
		4:  "Brentford",
		5:  "Brigton",
		6:  "Chelsea",
		7:  "Coventry",
		8:  "Crystal P.",
		9:  "Everton",
		10: "Fulham",
		11: "Hull C.",
		12: "Ipswitch T",
		13: "Leeds Utd",
		14: "Liverpool",
		15: "Man City",
		16: "Man Utd",
		17: "NewCastle",
		18: "Nottingham",
		19: "SunderLand",
		20: "TottenHam Hotspur",
	}

	return FixtureDict
}

func FindUpcomingMatches(games []fpl.Fixture) []fpl.Fixture {
	var output []fpl.Fixture
	i := 0
	for i < len(games) {
		if !(games[i].Finished) {
			output = append(output, games[i])
		}
		i = i + 1

	}
	return output
}

func FirstFive(games []fpl.Fixture) []fpl.Fixture {
	var output []fpl.Fixture
	i := 0
	if len(games) < 5 {
		for i < len(games) {
			output = append(output, games[i])
			i = i + 1
		}
		return output
	}

	for i < 5 {
		output = append(output, games[i])
		i = i + 1
	}
	return output
}

func FixturesForClub(team int, games []fpl.Fixture) []fpl.Fixture {
	var output []fpl.Fixture
	i := 0
	for i < len(games) {
		if games[i].TeamHome == team || games[i].TeamAway == team {
			output = append(output, games[i])
		}
		i = i + 1
	}
	return output
}

func GroupFirstFive(games []fpl.Fixture) map[int][]fpl.Fixture {
	output := map[int][]fpl.Fixture{}
	fixtures := FindUpcomingMatches(games)
	i := 1
	for i <= 5 {
		output[i] = FirstFive(FixturesForClub(i, fixtures))
		i = i + 1
	}
	return output
}

func FormatFixtures(clubs map[int][]fpl.Fixture, blocks map[int]string) string {
	output := ""

	for k, v := range clubs {
		var placeHolder string
		placeHolder += "Club " + strconv.Itoa(k)
		placeHolder += "\n"

		i := 0
		for i < len(v) {
			if v[i].TeamHome == k {
				placeHolder += " " + " vs " + strconv.Itoa(v[i].TeamAway) + " (H) " + "-" + "Difficulty " + blocks[v[i].TeamHomeDifficulty] + "\n"
			} else {
				placeHolder += " " + " vs " + strconv.Itoa(v[i].TeamHome) + " (A) " + "-" + "Difficulty " + blocks[v[i].TeamAwayDifficulty] + "\n"
			}
			i = i + 1
		}
		output = output + placeHolder

	}

	return output

}
