package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

// MCP Host (LLM Application)
// Will have multiple clients

func main() {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	messages := []api.Message{
		api.Message{
			Role:    "system",
			Content: "Provide very brief, concise responses",
		},
		api.Message{
			Role:    "user",
			Content: "Name some unusual animals",
		},
	}

	ctx := context.Background()
	req := &api.ChatRequest{
		Model:    "llama3.2",
		Messages: messages,
	}

	respFunc := func(resp api.ChatResponse) error {
		fmt.Print(resp.Message.Content)
		return nil
	}

	err = client.Chat(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
