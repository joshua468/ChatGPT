package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	apikey := os.Getenv("API_KEY")
	if apikey == "" {
		log.Fatalln("Missing Api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apikey)
	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"the first thing you should know in golang"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)

}
