package fpl

type Player struct {
	WebName     string `json:"web_name"`
	Position    int    `json:"element_type"`
	Club        int    `json:"team"`
	Cost        int    `json:"now_cost"`
	TotalPoints int    `json:"total_points"`

	GameweekPoints int `json:"event_points"`
}
