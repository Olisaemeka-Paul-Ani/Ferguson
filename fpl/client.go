package fpl

import (
	"encoding/json"
	"io"
	"net/http"
)

type Container struct {
	Group []Player `json:"elements"`
}

func FetchAllPlayers() ([]Player, error) {
	resp, err := http.Get("https://fantasy.premierleague.com/api/bootstrap-static/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var PlayerList Container
	err = json.Unmarshal(body, &PlayerList)

	if err != nil {
		return nil, err
	}

	return PlayerList.Group, nil
}
