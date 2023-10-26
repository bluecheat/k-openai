package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

func main() {
	config, err := LoadConfiguration("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	client := openai.NewClient(config.ApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "안녕반가워!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
