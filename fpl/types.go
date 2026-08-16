package fpl

type Player struct {
	WebName     string `json:"web_name"`
	Position    int    `json:"element_type"`
	Club        int    `json:"team"`
	Cost        int    `json:"now_cost"`
	TotalPoints int    `json:"total_points"`

	GameweekPoints int `json:"event_points"`
}

type Fixture struct {
	Event              int  `json:"event"`
	TeamHome           int  `json:"team_h"`
	TeamAway           int  `json:"team_a"`
	TeamHomeDifficulty int  `json:"team_h_difficulty"`
	TeamAwayDifficulty int  `json:"team_a_difficulty"`
	Finished           bool `json:"finished"`
}
