package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

// ImageGenerate https://platform.openai.com/docs/api-reference/images/create
func ImageGenerate(apiKey string, prompt string) (*openai.ImageResponse, error) {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			N:              1,
			ResponseFormat: openai.CreateImageResponseFormatB64JSON,
			Size:           openai.CreateImageSize256x256,
		},
	)

	if err != nil {
		fmt.Printf("ImageResponse error: %v\n", err)
		return nil, err
	}
	return &resp, nil
}
