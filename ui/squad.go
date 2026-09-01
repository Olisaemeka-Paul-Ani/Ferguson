package ui

import (
	"strconv"
	"strings"

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

func FormatData(squad [][]string) string {
	output := ""
	i := 0
	for i < len(squad) {
		placeHolder := squad[i]
		formatted := strings.Join(placeHolder, " ")
		output = output + formatted
		output = output + "\n"
		i = i + 1
	}
	return output
}

func GetPlayerDict() map[int]string {
	output := map[int]string{
		1: "GK",
		2: "DEF",
		3: "MID",
		4: "FOR",
	}
	return output
}
