package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Ollama struct {
	endPoint string
	model    string
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type chatResponse struct {
	Model              string      `json:"model"`
	CreatedAt          time.Time   `json:"created_at"`
	Message            chatMessage `json:"message"`
	DoneReason         string      `json:"done_reason"`
	Done               bool        `json:"done"`
	TotalDuration      int64       `json:"total_duration"`
	LoadDuration       int64       `json:"load_duration"`
	PromptEvalCount    int         `json:"prompt_eval_count"`
	PromptEvalDuration int64       `json:"prompt_eval_duration"`
	EvalCount          int         `json:"eval_count"`
	EvalDuration       int64       `json:"eval_duration"`
}

func NewOllama(model string) Ollama {
	return Ollama{
		model:    model,
		endPoint: "http://localhost:11434/api/chat",
	}
}

func (o Ollama) Generate(question string) (string, error) {
	reqBody := chatRequest{
		Model: o.model,
		Messages: []chatMessage{
			{"user", question},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(reqBodyBytes)
	req, err := http.NewRequest("POST", o.endPoint, reader)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("API request failed with status %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result chatResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		panic(err)
	}

	if len(result.Message.Content) > 0 {
		return result.Message.Content, nil
	}

	return "", fmt.Errorf("No answer received from the model.")
}
