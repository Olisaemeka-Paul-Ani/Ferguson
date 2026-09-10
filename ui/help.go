package ui

var Directions string = "DIRECTIONS"

var DirectionsCommands = map[string]string{
	"ENTER":      "View specific Player Stats",
	"ESC":        "Exit specific player stats",
	"UP/DOWN":    "Toggle layer selection/View depending on if enter command was pressed",
	"RIGHT/LEFT": "Toggle Active Pane",
	"?":          "Toggle Help/About Page",
}

var DirectionsCommandsKeys = []string{
	"?",
	"ENTER",
	"ESC",
	"RIGHT/LEFT",
	"UP/DOWN",
}
