package kopenai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

func (c *NaverOpenApi) Transition(ctx context.Context, request *TranslationRequest) (response *TranslationResponse, err error) {
	apiUrl := naverOpenApi + papagoTranslation
	req, err := c.createRequest(ctx, http.MethodPost, apiUrl, request)
	if err != nil {
		return
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if isFailStatusCode(res) {
		var errRes ErrorPapaoResponse
		err = decodeResponse(res.Body, &errRes)
		err = errors.New(fmt.Sprintf("[%s]%s", errRes.ErrorCode, errRes.ErrorMessage))
		return
	}
	err = decodeResponse(res.Body, &response)
	return
}

type Lang string

const (
	KO   Lang = "ko"
	EN   Lang = "en"
	JA   Lang = "ja"
	ZHCN Lang = "zh-CN"
	ZHTW Lang = "zh-TW"
	VI   Lang = "vi"
	ID   Lang = "id"
	TH   Lang = "th"
	DE   Lang = "de"
	RU   Lang = "ru"
	ES   Lang = "es"
	IT   Lang = "it"
	FR   Lang = "fr"
)

func (l Lang) String() string {
	return string(l)
}

type TranslationRequest struct {
	Source Lang   `json:"source"`
	Target Lang   `json:"target"`
	Text   string `json:"text"`
}

type TranslationResponse struct {
	Message struct {
		Type    string `json:"@type"`
		Service string `json:"@service"`
		Version string `json:"@version"`
		Result  TranslationResult
	} `json:"message"`
}

type TranslationResult struct {
	SrcLangType    string `json:"srcLangType"`
	TarLangType    string `json:"tarLangType"`
	TranslatedText string `json:"translatedText"`
}

type ErrorPapaoResponse struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
