package ui

import (
	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
)

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
	for i <= 20 {
		output[i] = FirstFive(FixturesForClub(i, fixtures))
		i = i + 1
	}
	return output
}
