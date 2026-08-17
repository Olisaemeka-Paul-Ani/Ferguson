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
