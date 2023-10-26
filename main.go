package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-openai-gpt/gpt"
	"log"
	"os"
)

func main() {
	config, err := LoadConfiguration("env.yaml")
	if err != nil {
		log.Fatal(err)
	}

	resp2, err := gpt.ImageGenerate(config.ApiKey, "이별문자 생성기에 대한 단순한 이미지를 생성해줘")
	if err != nil {
		return
	}
	prettyPrint(resp2)

	// Open output file
	dec, err := base64.StdEncoding.DecodeString(resp2.Data[0].B64JSON)
	if err != nil {
		panic(err)
	}

	output, err := os.Create("created-image.png")
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

func prettyPrint(obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
