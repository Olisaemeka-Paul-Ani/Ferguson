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

func GetForm(PrevPoints []fpl.Points, ColorMap map[int]string) string {
	count := 5
	output := ""
	i := len(PrevPoints) - 1
	for count > 0 && i >= 0 {
		if PrevPoints[i].PointsForPerson < 0 {
			output = output + ColorMap[-1]
		} else if PrevPoints[i].PointsForPerson > 9 {
			output = output + ColorMap[10]
		} else {
			output = output + ColorMap[PrevPoints[i].PointsForPerson]

		}
		count = count - 1
		i = i - 1
	}
	return output

}

var StatusMap = map[string]string{
	"a": "Available",
	"d": "Doubtful",
	"i": "Injured",
	"s": "Suspended",
	"u": "Unavailable",
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
