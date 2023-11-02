# k-openai

k-openai는 Go 언어로 작성된 오픈소스 프로젝트로, OpenAI와 Naver Open API를 사용하여 다양한 기능을 제공합니다.

## 설치 방법

1. Go가 설치되어 있지 않다면, [Go 공식 웹사이트](https://golang.org/dl/)에서 Go를 설치하세요.
2. 터미널 또는 명령 프롬프트를 열고 다음 명령어를 실행하여 이 프로젝트를 클론하세요:
```
    go get github.com/bluecheat/k-openai
    go mod tidy
```
3. 다음 명령어를 실행하여 필요한 Go 모듈을 설치하세요:

## 사용 방법
1. main.go 파일을 생성하고 다음과 같은 기본 코드를 작성하세요:
```go
package main

import (
    "fmt"
    "log"
    "github.com/bluecheat/k-openai"
)

func main() {
    // 환경 설정 로드
    config, err := kopenai.LoadConfiguration()
    if err != nil {
        log.Fatal(err)
    }

    // OpenAI 클라이언트 초기화
    openaiClient := kopenai.NewOpenAIClient(config.OpenAI.ApiKey)

    // 예제: OpenAI를 사용하여 텍스트 생성
    response, err := openaiClient.CreateCompletion("Hello, world! How are you today?")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("OpenAI Response:", response)

    // Naver 클라이언트 초기화
    naverClient := kopenai.NewNaverClient(config.Naver.ClientID, config.Naver
```