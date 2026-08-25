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
