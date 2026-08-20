package ai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

// Marshalling response to send to Gemini API
type Text struct {
	Text string `json:"text"`
}

type Part struct {
	Parts []Text `json:"parts"`
}

type Outgoing struct {
	Content []Part `json:"contents"`
}

// Unmarshalling response from Gemini API
type Response struct {
	ResponseText string `json:"text"`
}

type PartsStruct struct {
	PartText []Response `json:"parts"`
}

type ContentStruct struct {
	Content PartsStruct `json:"content"`
}

type Candidates struct {
	Candidate []ContentStruct `json:"candidates"`
}

func AggregatePrompt(str string) ([]byte, error) {
	var prompt = Text{"Say Hello in one word, in any niche language of your choice."}
	var TextArray []Text
	TextArray = append(TextArray, prompt)
	var PartStruct = Part{TextArray}
	var MessageArray []Part
	MessageArray = append(MessageArray, PartStruct)

	var Message = Outgoing{MessageArray}
	JsonBytes, err := json.Marshal(Message)
	if err != nil {
		return nil, err
	}

	return JsonBytes, nil
}

func SendPrompt(byt []byte) ([]byte, error) {
	resp, err := http.Post("https://generativelanguage.googleapis.com/v1beta/models/gemini-3.6-flash:generateContent?key="+os.Getenv("FERGUSON_AI_KEY"), "application/json", bytes.NewBuffer(byt))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
