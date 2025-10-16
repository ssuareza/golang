package vision

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	visionURL = "https://api.openai.com/v1/chat/completions"
)

// Request represents the structure for the request to the OpenAI Vision API.
type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message represents the structure for a message in the request.
type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

// MessageResponse represents the structure of the message within a choice in the response.
type MessageResponse struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Content represents the structure for content within a message.
type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

// ImageURL represents the structure for an image URL in the request.
type ImageURL struct {
	URL string `json:"url"`
}

// Response represents the structure of the response from the OpenAI Vision API.
type Response struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Choice represents the structure for a choice in the response.
type Choice struct {
	Index        int             `json:"index"`
	Message      MessageResponse `json:"message"`
	Logprobs     *string         `json:"logprobs"`
	FinishReason string          `json:"finish_reason"`
}

// Usage represents the structure for the usage statistics in the response.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// SendImage sends an image and a prompt to the OpenAI Vision API and returns the response.
func SendImage(ctx context.Context, apiKey string, imagePath string, prompt string) (*Response, error) {
	// Read image
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// Post image to visionURL
	reqBody := Request{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: prompt,
					},
					{
						Type: "image_url",
						ImageURL: &ImageURL{
							URL: fmt.Sprintf("data:image/jpeg;base64,%s", base64Image),
						},
					},
				},
			},
		},
	}

	// Marshal the request body to JSON
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", visionURL, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer res.Body.Close()

	// Read the response body
	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for non-200 status codes
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-OK status: %d - %s", res.StatusCode, string(resBodyBytes))
	}

	// Unmarshal the response body
	var r Response
	err = json.Unmarshal(resBodyBytes, &r)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &r, nil
}
