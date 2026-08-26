package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

// GEMINI API STRUCTS, REQUEST, RESPONSE WORK (BEGINNING)
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
	var prompt = Text{str}
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

func UnmarshallResponseBody(byt []byte) (Candidates, error) {
	var ResponseStruct Candidates
	err := json.Unmarshal(byt, &ResponseStruct)
	if err != nil {
		return ResponseStruct, err
	}

	return ResponseStruct, nil
}

func ExtractunMarshalledResponse(response Candidates) (string, error) {
	var err error
	if len(response.Candidate) == 0 {
		err = errors.New("No content inside response")
		return "", err
	}

	if len(response.Candidate[0].Content.PartText) == 0 {
		err = errors.New("No content inside response")
		return "", err
	}

	return response.Candidate[0].Content.PartText[0].ResponseText, err
}

// GEMINI API STRUCTS, REQUEST, RESPONSE WORK (ENDING)

// GROQ API STRUCTS, REQUEST, RESPONSE WORK (BEGINNING)
// Marshalling request to send to Groq API
type Body struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model    string `json:"model"`
	Messages []Body `json:"messages"`
}

type GroqResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqResponseChoicesContent struct {
	Index   int                 `json:"index"`
	Message GroqResponseMessage `json:"message"`
}

type GroqResponseContent struct {
	Choices []GroqResponseChoicesContent `json:"choices"`
}

func AggregateGroqRequest(str string) ([]byte, error) {
	var prompt = Body{"user", str}
	var messagesSlice []Body
	messagesSlice = append(messagesSlice, prompt)
	var request = GroqRequest{"qwen/qwen3.6-27b", messagesSlice}
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func SendGroqRequest(byt []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(byt))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("FERGUSON_GROQ_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

func UnMarshallGroqResponse(byt []byte) (GroqResponseContent, error) {
	var response GroqResponseContent
	err := json.Unmarshal(byt, &response)
	if err != nil {
		return response, err
	}

	return response, nil

}

func ExtractUnmarshalledGrokResponse(res GroqResponseContent) (string, error) {
	var err error
	if len(res.Choices) == 0 {
		err = errors.New("No content inside response")
		return "", err
	}

	return res.Choices[0].Message.Content, err
}

// GROQ API STRUCTS, REQUEST, RESPONSE WORK (ENDING)

// LOGIC FOR WIRING PROMPT FROM PROMPT.GO INTO BOTH LLMS  (BEGINNING)
func GetGrokReply() (string, error) {

	PromptStr, err := ConvertBundledData()
	if err != nil {
		return "", err
	}

	var byt []byte

	byt, err = AggregateGroqRequest(PromptStr)

	if err != nil {
		return "", err
	}

	byt, err = SendGroqRequest(byt)

	if err != nil {
		return "", err
	}
	var GroqResStruct GroqResponseContent
	GroqResStruct, err = UnMarshallGroqResponse(byt)
	if err != nil {
		return "", err
	}

	var GroqResString string
	GroqResString, err = ExtractUnmarshalledGrokResponse(GroqResStruct)

	if err != nil {
		return "", err
	}

	return GroqResString, nil
}

func GetGeminiReply() (string, error) {
	PromptStr, err := ConvertBundledData()
	if err != nil {
		return "", err
	}

	var byt []byte

	byt, err = AggregatePrompt(PromptStr)
	if err != nil {
		return "", err
	}

	byt, err = SendPrompt(byt)

	if err != nil {
		return "", err
	}

	var GeminiRespStruct Candidates

	GeminiRespStruct, err = UnmarshallResponseBody(byt)

	if err != nil {
		return "", err
	}

	var output string

	output, err = ExtractunMarshalledResponse(GeminiRespStruct)

	if err != nil {
		return "", err
	}

	return output, nil

}

//LOGIC FOR WIRING PROMPT FROM PROMPT.GO INTO BOTH LLMS  (BEGINNING)
