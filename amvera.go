package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// AmveraClient структура клиента
type AmveraClient struct {
	apiBase       string // базовый URL (например, https://llm.api.amvera.io)
	inferenceName string // имя инференса: gpt, llama, deepseek, qwen
	token         string
	model         string // конкретная модель: gpt-4.1, gpt-5, llama8b и т.д.
	systemPrompt  string
	httpClient    *http.Client
}

// NewAmveraClient создаёт нового клиента из переменных окружения
func NewAmveraClient() *AmveraClient {
	return &AmveraClient{
		apiBase:       getEnv("AMVERA_API_BASE", "https://llm.api.amvera.io"),
		inferenceName: getEnv("AMVERA_INFERENCE_NAME", "gpt"),
		token:         os.Getenv("AMVERA_TOKEN"),
		model:         getEnv("AMVERA_MODEL", "gpt-5"),
		systemPrompt:  os.Getenv("AMVERA_SYSTEM_PROMPT"),
		httpClient:    &http.Client{},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type ResponseBody struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Response string `json:"response,omitempty"`
}

func (c *AmveraClient) Ask(prompt string) (string, error) {

	if c.token == "" {
		return "", fmt.Errorf("AMVERA_TOKEN is not set")
	}

	messages := []Message{}

	if c.systemPrompt != "" {
		messages = append(messages, Message{Role: "system", Text: c.systemPrompt})
	}

	messages = append(messages, Message{Role: "user", Text: prompt})

	reqBody := RequestBody{
		Model:    c.model,
		Messages: messages,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/models/%s", c.apiBase, c.inferenceName)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request URL: %s", url)
		log.Printf("Request body: %s", string(jsonData))
		return "", fmt.Errorf("API returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	var respBody ResponseBody
	if err := json.Unmarshal(body, &respBody); err != nil {
		return "", fmt.Errorf("failed to parse response JSON: %w, body: %s", err, string(body))
	}

	if len(respBody.Choices) > 0 && respBody.Choices[0].Message.Content != "" {
		return respBody.Choices[0].Message.Content, nil
	}
	if respBody.Response != "" {
		return respBody.Response, nil
	}

	return string(body), nil
}
