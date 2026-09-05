package fpl

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

func FetchAllFixtures() ([]Fixture, error) {
	resp, err := http.Get("https://fantasy.premierleague.com/api/fixtures/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var FixtureList []Fixture
	err = json.Unmarshal(body, &FixtureList)

	if err != nil {
		return nil, err
	}

	return FixtureList, nil

}

func FetchFormData(id int) ([]Points, error) {
	enp := "https://fantasy.premierleague.com/api/element-summary/"
	enp = enp + strconv.Itoa(id) + "/"
	resp, err := http.Get(enp)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var FormList FormStruct
	err = json.Unmarshal(body, &FormList)

	if err != nil {
		return nil, err
	}

	return FormList.Form, nil
}
