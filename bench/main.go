package main

import (
	"fmt"
	"time"

	"github.com/Olisaemeka-Paul-Ani/ferguson/fpl"
)

func main() {
	start := time.Now()
	ch := make(chan time.Duration)

	go func() {
		start := time.Now()

		_, _ = fpl.FetchAllPlayers()

		latency := time.Since(start)
		ch <- latency

	}()

	go func() {
		start := time.Now()

		_, _ = fpl.FetchAllFixtures()

		latency := time.Since(start)
		ch <- latency
	}()

	go func() {
		start := time.Now()

		_, _ = fpl.FetchSquadPlayers(565066)

		latency := time.Since(start)
		ch <- latency
	}()

	i := 0
	var result time.Duration
	for i < 3 {
		result = result + <-ch
		i = i + 1

	}
	fmt.Println(result)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
