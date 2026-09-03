package fpl

type Player struct {
	WebName     string `json:"web_name"`
	Position    int    `json:"element_type"`
	Club        int    `json:"team"`
	Cost        int    `json:"now_cost"`
	TotalPoints int    `json:"total_points"`

	GameweekPoints  int    `json:"event_points"`
	News            string `json:"news"`
	Status          string `json:"status"`
	ChanceOfPlaying *int   `json:"chance_of_playing_next_round"`
}

type Fixture struct {
	Event              int  `json:"event"`
	TeamHome           int  `json:"team_h"`
	TeamAway           int  `json:"team_a"`
	TeamHomeDifficulty int  `json:"team_h_difficulty"`
	TeamAwayDifficulty int  `json:"team_a_difficulty"`
	Finished           bool `json:"finished"`
}
