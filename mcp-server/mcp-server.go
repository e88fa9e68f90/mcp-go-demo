package mcp_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func MCPStreamableServer() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo",
		"1.0.0",
	)
	sseServer := server.NewStreamableHTTPServer(s, server.WithEndpointPath("/mcp"))
	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)
	// Add tool handler
	s.AddTool(tool, helloHandler)
	s.AddPrompt(mcp.NewPrompt("greeting",
		mcp.WithPromptDescription("一个友好的问候提示"),
		mcp.WithArgument("name",
			mcp.ArgumentDescription("要问候的人的名字"),
		),
	), func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		name := request.Params.Arguments["name"]
		if name == "" {
			name = "朋友"
		}

		return mcp.NewGetPromptResult(
			"友好的问候",
			[]mcp.PromptMessage{
				mcp.NewPromptMessage(
					mcp.RoleAssistant,
					mcp.NewTextContent(fmt.Sprintf("你好，%s！今天有什么可以帮您的吗？", name)),
				),
			},
		), nil
	})

	// Start the stdio server
	if err := sseServer.Start(":8000"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
func MCPSSEServer() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo",
		"1.0.0",
	)
	sseServer := server.NewSSEServer(s)
	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)
	// Add tool handler
	s.AddTool(tool, helloHandler)

	// Start the stdio server
	if err := sseServer.Start(":8000"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return nil, errors.New("name must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
