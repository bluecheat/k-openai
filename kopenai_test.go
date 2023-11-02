package kopenaigpt

import (
	"context"
	"encoding/base64"
	"github.com/sashabaranov/go-openai"
	"os"
	"testing"
	"time"
)

func TestKopenai_Chat(t *testing.T) {

	config, err := LoadConfiguration("env.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	kopenai := NewKopenAiGpt(config)

	ctx := context.Background()
	resp, err := kopenai.Chat(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "안녕 너의 이름은 뭐니?",
			},
		},
	}, ChatTransOption{
		InputPrompt: &TransOption{
			Source: KO,
			Target: EN,
		},
		OutputPrompt: &TransOption{
			Source: EN,
			Target: KO,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.Choices[0].Message)
}

func TestKopenai_Image(t *testing.T) {

	config, err := LoadConfiguration("env.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	kopenai := NewKopenAiGpt(config)

	ctx := context.Background()
	resp, err := kopenai.ImageGenerate(ctx, openai.ImageRequest{
		Prompt:         "업무, 개발, AI개발",
		N:              1,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}, ImageTransOption{
		InputPrompt: &TransOption{
			Source: KO,
			Target: EN,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	// Open output file
	dec, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		panic(err)
	}
	output, err := os.Create("created-image-" + time.Now().UTC().String() + ".png")
	if err != nil {
		panic(err)
	}
	// Close output file
	defer output.Close()

	if _, err := output.Write(dec); err != nil {
		panic(err)
	}
	if err := output.Sync(); err != nil {
		panic(err)
	}
}
