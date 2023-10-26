package main

import (
	"encoding/json"
	"fmt"
	"go-openapi-gp/gpt"
	"log"
)

func main() {
	config, err := LoadConfiguration("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := gpt.Chat(config.ApiKey, "내일 날씨에 대해서 예상해줘")
	if err != nil {
		return
	}
	prettyPrint(resp)

	resp2, err := gpt.ImageGenerate(config.ApiKey, "이별문자 생성기에 대한 단순한 이미지를 생성해줘")
	if err != nil {
		return
	}
	prettyPrint(resp2)
}

func prettyPrint(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
