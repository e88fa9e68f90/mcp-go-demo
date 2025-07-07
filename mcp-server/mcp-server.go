package mcp_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func MCPServer() {
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
	calculatorTool := mcp.NewTool("calculate",
		mcp.WithDescription("执行基本的算术运算"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("要执行的算术运算类型"),
			mcp.Enum("add", "subtract", "multiply", "divide"), // 保持英文
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("第一个数字"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("第二个数字"),
		),
	)
	// Add tool handler
	s.AddTool(tool, helloHandler1)
	s.AddTool(calculatorTool, calculatorTool2)

	// Start the stdio server
	if err := sseServer.Start(":8000"); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func helloHandler1(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func calculatorTool2(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	op := request.Params.Arguments["operation"].(string)
	x := request.Params.Arguments["x"].(float64)
	y := request.Params.Arguments["y"].(float64)

	var result float64
	switch op {
	case "add":
		result = x + y
	case "subtract":
		result = x - y
	case "multiply":
		result = x * y
	case "divide":
		if y == 0 {
			return nil, errors.New("不允许除以零")
		}
		result = x / y
	}

	return mcp.FormatNumberResult(result), nil
}
