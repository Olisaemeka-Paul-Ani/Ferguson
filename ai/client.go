package ai

type Text struct {
	Text string `json:"text"`
}

type Part struct {
	Parts []Text `json:"parts"`
}

type Outgoing struct {
	Content []Part `json:"contents"`
}
