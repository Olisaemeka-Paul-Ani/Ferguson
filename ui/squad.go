package ui

import (
	"strconv"

	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
)

func CleanData(squad []fpl.Player) [][]string {
	output := [][]string{}
	i := 0
	for i < len(squad) {
		placeHolder := []string{}
		placeHolder = append(placeHolder, squad[i].WebName)
		placeHolder = append(placeHolder, strconv.Itoa(squad[i].Position))
		placeHolder = append(placeHolder, strconv.Itoa(squad[i].Club))
		placeHolder = append(placeHolder, strconv.Itoa(squad[i].Cost))
		placeHolder = append(placeHolder, strconv.Itoa(squad[i].TotalPoints))
		placeHolder = append(placeHolder, strconv.Itoa(squad[i].GameweekPoints))
		output = append(output, placeHolder)
		i = i + 1

	}
	return output

}
