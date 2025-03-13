package main

import "github.com/mark3labs/mcp-go/client"

func main() {
	c, err := client.NewStdioMCPClient("go", []string{}, "run", "../server/main.go")

}
