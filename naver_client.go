package kopenai

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	naverOpenApi      = "https://openapi.naver.com"
	papagoTranslation = "/v1/papago/n2mt"
)

type NaverOpenApi struct {
	config         NaverOpenApiConfig
	requestBuilder *RequestBuilder
	httpClient     *http.Client
}

func NewNaverOpenApiClient(config NaverOpenApiConfig) *NaverOpenApi {
	return &NaverOpenApi{
		config:         config,
		requestBuilder: NewRequestBuilder(),
		httpClient:     &http.Client{},
	}
}

type RequestBuilder struct{}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

func (b *RequestBuilder) Build(
	ctx context.Context,
	method string,
	url string,
	body any,
	header http.Header,
) (req *http.Request, err error) {
	var bodyReader io.Reader
	if body != nil {
		if v, ok := body.(io.Reader); ok {
			bodyReader = v
		} else {
			var reqBytes []byte
			reqBytes, err = json.Marshal(body)
			if err != nil {
				return
			}
			bodyReader = bytes.NewBuffer(reqBytes)
		}
	}
	req, err = http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return
	}
	if header != nil {
		req.Header = header
	}
	return
}

func (c *NaverOpenApi) createRequest(ctx context.Context, method string, url string, body any) (*http.Request, error) {
	req, err := c.requestBuilder.Build(ctx, method, url, withUrlEncodeBody(body), nil)
	if err != nil {
		return nil, err
	}
	c.setRequestHeaders(req)

	return req, nil
}

func (c *NaverOpenApi) setRequestHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("X-Naver-Client-Id", c.config.ClientId)
	req.Header.Set("X-Naver-Client-Secret", c.config.ClientSecret)
}

func isFailStatusCode(resp *http.Response) bool {
	return resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}
	return json.NewDecoder(body).Decode(v)
}

func withJsonBody(body any) (io.Reader, error) {
	var reqBytes []byte
	reqBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(reqBytes), nil
}

func withUrlEncodeBody(body any) io.Reader {
	return bytes.NewBufferString(StructToUrlValues(body).Encode())
}
