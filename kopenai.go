package kopenaigpt

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"log"
)

var (
	DefaultChatRequest = openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
			},
		},
	}

	DefaultImageRequest = openai.ImageRequest{
		N:              1,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		Size:           openai.CreateImageSize512x512,
	}
)

type KopenAiGpt struct {
	naverClient  *NaverOpenApi
	openaiClient *openai.Client

	config *AiConfig
}

func NewKopenAiGpt(config *AiConfig) *KopenAiGpt {
	validConfig(config)
	client := openai.NewClient(config.Openai.ApiKey)
	return &KopenAiGpt{
		config:       config,
		openaiClient: client,
		naverClient:  NewNaverOpenApiClient(config.Naver),
	}
}

type ChatTransOption struct {
	InputPrompt  *TransOption
	OutputPrompt *TransOption
}

type ImageTransOption struct {
	InputPrompt *TransOption
}

type TransOption struct {
	Source Lang
	Target Lang
}

// Chat https://platform.openai.com/docs/api-reference/chat
func (k *KopenAiGpt) Chat(ctx context.Context, request openai.ChatCompletionRequest, option ChatTransOption) (resp openai.ChatCompletionResponse, err error) {
	err = k.transChatRequest(ctx, &request, option)
	if err != nil {
		//fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	resp, err = k.openaiClient.CreateChatCompletion(
		ctx,
		request,
	)
	if err != nil {
		//fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	err = k.transChatResponse(ctx, &resp, option)
	return
}

// ImageGenerate https://platform.openai.com/docs/api-reference/images/create
func (k *KopenAiGpt) ImageGenerate(ctx context.Context, request openai.ImageRequest, option ImageTransOption) (resp openai.ImageResponse, err error) {
	err = k.transImageRequest(ctx, &request, option)

	resp, err = k.openaiClient.CreateImage(
		ctx,
		request,
	)

	if err != nil {
		//fmt.Printf("ImageGenerate error: %v\n", err)
		return
	}
	return
}

func (k *KopenAiGpt) transChatRequest(ctx context.Context, request *openai.ChatCompletionRequest, option ChatTransOption) error {
	for i, text := range request.Messages {
		var prompt string
		if option.InputPrompt != nil {
			transText, err := k.naverClient.Transition(ctx, &TranslationRequest{
				Source: option.InputPrompt.Source,
				Target: option.InputPrompt.Target,
				Text:   text.Content,
			})
			if err != nil {
				return err
			}
			prompt = transText.Message.Result.TranslatedText
		} else {
			prompt = text.Content
		}
		request.Messages[i].Content = prompt
	}
	return nil
}

func (k *KopenAiGpt) transChatResponse(ctx context.Context, response *openai.ChatCompletionResponse, option ChatTransOption) error {
	for i, result := range response.Choices {
		if option.OutputPrompt != nil {
			transText, err := k.naverClient.Transition(ctx, &TranslationRequest{
				Source: option.OutputPrompt.Source,
				Target: option.OutputPrompt.Target,
				Text:   result.Message.Content,
			})
			if err != nil {
				return err
			}
			response.Choices[i].Message.Content = transText.Message.Result.TranslatedText
		}
	}
	return nil
}

func (k *KopenAiGpt) transImageRequest(ctx context.Context, request *openai.ImageRequest, option ImageTransOption) error {
	if option.InputPrompt != nil {
		transText, err := k.naverClient.Transition(ctx, &TranslationRequest{
			Source: option.InputPrompt.Source,
			Target: option.InputPrompt.Target,
			Text:   request.Prompt,
		})
		if err != nil {
			return err
		}
		request.Prompt = transText.Message.Result.TranslatedText
	}
	return nil
}

func validConfig(config *AiConfig) {
	if config == nil || config.Openai.ApiKey == "" || config.Naver.ClientId == "" || config.Naver.ClientSecret == "" {
		log.Fatalln("empty openai key or naver client info")
	}
}
