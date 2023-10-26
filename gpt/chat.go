package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// Chat https://platform.openai.com/docs/api-reference/chat
func Chat(apiKey string, prompt string) (*openai.ChatCompletionResponse, error) {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	return &resp, nil
}
