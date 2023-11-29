# K-Openai

K-openai는 OpenAI의 API 사용을 위한 프롬프트 작성 시 파파고 번역기능을 통해 Token 사용량 최소화와 서비스 형태에 맞는 응답을 하기 위해 제작하게 되었습니다.

## 목차
- [1. 사전 준비](#사전-준비)
- [2. 설치 방법](#설치-방법)
- [3. Config 설정 방법](#Config-설정-방법)
- [4. 사용 방법](#사용-방법)
  - Chat API (GPT 채팅 API)
  - ImageGenerate API ( 이미지 생성 API )

## 사전 준비
- `go version >= 1.18` 이상 설치 되어 있어야합니다.
- [OpenAI](https://platform.openai.com/) 에서 API Key를 발급 받아야합니다.
- [Naver Developers](https://developers.naver.com/main/)에서 client를 발급 받아야합니다.


## 설치 방법
터미널 또는 명령 프롬프트를 열고 다음 명령어를 실행하여 의존성을 설치합니다.
```
go get -u github.com/bluecheat/k-openai@v0.1.3
```

## Config 설정 방법

#### 1. yaml 파일을 읽어서 설정하는 방법
```go
config, err := kopenai.LoadConfiguration("./kopenai_env.yml")
if err != nil {
    log.Fatal(err)
}
```

#### 2. 변수 직접 선언 방법
```go
config := kopenai.Config{
    Openai: kopenai.OpenAiConfig{
        ApiKey: "apikey",
    },
    Naver: kopenai.NaverOpenApiConfig{
        ClientId:     "clinetId",
        ClientSecret: "clientSecret",
    },
}
```

## 사용 방법

### 1. Chat API `kopenai.Chat` ( GPT 채팅 API )
[Chat API](https://platform.openai.com/docs/api-reference/chat)에서 사용하는 요청값, 응답값은 동일하고 추가적으로 번역 옵션을 넣어서 처리합니다.

#### 파라미터 정보
  - `openai.ChatCompletionRequest`: 기존의 openai의 Chat 호출 시 사용되는 요청값
  - `kopenai.ChatTransOption`: 프롬프트 번역 시 사용하는 Option
    - InputPrompt: 프롬프트에 대한 번역 ( source: 기존 언어, target: 번역할 언어 )
    - OutputPrompt: 프롬프트 결과에 대한 번역 ( source: 기존 언어, target: 번역할 언어 )

####  예제 코드
```go
...

client := kopenai.NewKopenAiGpt(config)

ctx := context.Background()
resp, err := client.Chat(ctx, openai.ChatCompletionRequest{
    Model: openai.GPT3Dot5Turbo,
    Messages: []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleUser,
            Content: "안녕 너의 이름은 뭐니?",
        },
    },
}, kopenai.ChatTransOption{
    InputPrompt: &kopenai.TransOption{
        Source: kopenai.KO,
        Target: kopenai.EN,
    },
    OutputPrompt: &kopenai.TransOption{
        Source: kopenai.EN,
        Target: kopenai.KO,
    },
})
if err != nil {
    t.Error(err)
    return
}
fmt.Println(resp.Choices[0].Message)
```

### 2. ImageGenerate API `kopenai.ImageGenerate` ( 이미지 생성 API )
[ImageGenerate API](https://platform.openai.com/docs/api-reference/images/create)에서 사용하는 요청값, 응답값은 동일하고 추가적으로 번역 옵션을 넣어서 처리합니다.
#### 파라미터 정보
- `openai.ImageRequest`: 기존의 openai의 ImageGenerate 호출 시 사용되는 요청값
- `kopenai.ImageTransOption`: 프롬프트 번역 시 사용하는 Option
    - InputPrompt: 프롬프트에 대한 번역 ( source: 기존 언어, target: 번역할 언어 )

####  예제 코드

```go
...
client := kopenai.NewKopenAiGpt(config)

ctx := context.Background()
resp, err := client.ImageGenerate(ctx, openai.ImageRequest{
    Prompt:         "업무, 개발, AI개발",
    N:              1,
    Size:           openai.CreateImageSize256x256,
    ResponseFormat: openai.CreateImageResponseFormatB64JSON,
}, kopenai.ImageTransOption{
    InputPrompt: &kopenai.TransOption{
        Source: kopenai.KO,
        Target: kopenai.EN,
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
```

#### 이미지 생성 결과
![image](https://github.com/bluecheat/k-openai/assets/55500108/36b1593b-37d3-4694-830a-5304e0a742c0)
