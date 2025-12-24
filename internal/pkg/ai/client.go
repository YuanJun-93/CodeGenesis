package ai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/YuanJun-93/CodeGenesis/internal/config"
	"github.com/YuanJun-93/CodeGenesis/internal/pkg/constant"
)

type AiClient struct {
	ApiKey  string
	BaseUrl string
	Model   string
	Client  *http.Client
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type StreamResponse struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

const (
// Removed local constants in favor of shared package
)

var defaultBaseUrls = map[string]string{
	constant.AiProviderOpenAI:   "https://api.openai.com/v1",
	constant.AiProviderDeepSeek: "https://api.deepseek.com",
	constant.AiProviderQwen:     "https://dashscope.aliyuncs.com/compatible-mode/v1",
}

var defaultModels = map[string]string{
	constant.AiProviderOpenAI:   "gpt-3.5-turbo",
	constant.AiProviderDeepSeek: "deepseek-chat",
	constant.AiProviderQwen:     "qwen-turbo",
}

func NewAiClient(c config.Config) *AiClient {
	baseUrl := c.Ai.BaseUrl
	model := c.Ai.Model
	provider := strings.ToLower(c.Ai.Provider)

	// If BaseUrl is missing, try to use Provider default
	if baseUrl == "" {
		if url, ok := defaultBaseUrls[provider]; ok {
			baseUrl = url
		} else {
			// Fallback default
			baseUrl = defaultBaseUrls[constant.AiProviderOpenAI]
		}
	}

	// If Model is missing, try to use Provider default
	if model == "" {
		if m, ok := defaultModels[provider]; ok {
			model = m
		} else {
			// Fallback default
			model = defaultModels[constant.AiProviderOpenAI]
		}
	}

	return &AiClient{
		ApiKey:  c.Ai.ApiKey,
		BaseUrl: baseUrl,
		Model:   model,
		Client:  &http.Client{},
	}
}

// DoRequest performs a synchronous chat completion request
func (c *AiClient) DoRequest(systemPrompt, userPrompt string) (string, error) {
	reqBody := ChatRequest{
		Model: c.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Stream: false,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", c.BaseUrl+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI API returned status: %d", resp.StatusCode)
	}

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", err
	}

	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}
	return "", fmt.Errorf("no response from AI")
}

// DoStreamRequest performs a streaming chat completion request
func (c *AiClient) DoStreamRequest(systemPrompt, userPrompt string) (chan string, error) {
	reqBody := ChatRequest{
		Model: c.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt + "\nPlease reply in Markdown format code block."},
			{Role: "user", Content: userPrompt},
		},
		Stream: true,
	}

	jsonBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", c.BaseUrl+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("AI API returned status: %d", resp.StatusCode)
	}

	stream := make(chan string)
	go func() {
		defer close(stream)
		defer resp.Body.Close()

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Printf("Stream read error: %v\n", err)
				}
				break
			}

			line = strings.TrimSpace(line)
			if !strings.HasPrefix(line, "data: ") {
				continue
			}

			data := strings.TrimPrefix(line, "data: ")
			if data == constant.StreamDone {
				break
			}

			var streamResp StreamResponse
			if err := json.Unmarshal([]byte(data), &streamResp); err != nil {
				continue
			}

			if len(streamResp.Choices) > 0 {
				content := streamResp.Choices[0].Delta.Content
				if content != "" {
					stream <- content
				}
			}
		}
	}()

	return stream, nil
}
