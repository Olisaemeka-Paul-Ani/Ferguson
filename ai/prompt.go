package ai

import (
	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
	"github.com/Olisaemeka-Paul-Ani/ferguson/ui"
)

func BundlePlayerData(strInput string) (string, error) {
	test, err := fpl.FetchAllPlayers()
	if err != nil {
		return "", err
	}

	str := ui.CleanData(test)
	output := ui.FormatData(str)

	strInput = strInput + output

	return strInput, nil

}

func BundleFixtureData(strInput string) (string, error) {
	resp, err := fpl.FetchAllFixtures()
	if err != nil {
		return "", err
	}

	var outputHash map[int][]fpl.Fixture
	var difficulty = map[int]string{
		1: "Very Easy",
		2: "Easy",
		3: "Neutral: Not easy, Not too hard",
		4: "Hard",
		5: "Very Hard",
	}

	outputHash = ui.GroupFirstFive(resp)
	output := ui.FormatFixtures(outputHash, difficulty)
	strInput = strInput + output
	return strInput, nil

}
