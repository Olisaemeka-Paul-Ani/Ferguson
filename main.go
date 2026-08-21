package main

import (
	"flag"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	nFlag := flag.Int("team", 0, "This is supposed to be your FPL team ID")
	flag.Parse()
	fmt.Println(*nFlag)
	m := Model{}
	p := tea.NewProgram(m)

	_, err := p.Run()

	if err != nil {
		fmt.Println("error:", err)
	}

}
