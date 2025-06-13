package mcp

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)
	// Add tool
	tool := mcp.NewTool("current time",
		mcp.WithDescription("Get current time with timezone, Asia/Shanghai is default"),
		mcp.WithString("timezone",
			mcp.Required(),
			mcp.Description("current time timezone"),
		),
	)
	// Add tool handler
	s.AddTool(tool, currentTimeHandler)

	// // Add tool
	// tool = mcp.NewTool("ep-copilot",
	// 	mcp.WithDescription("介绍ep-copilot的功能以及作用"),
	// 	// ),
	// )
	// // Add tool handler
	// s.AddTool(tool, epcopilotDescriptionHandler)
	// // 添加新的Git提交工具
	// tool = mcp.NewTool("提交代码",
	// 	mcp.WithDescription("在当前打开的工程目录里检查Go文件规范，并提交代码"),
	// 	mcp.WithString("message",
	// 		mcp.Description("提交信息"),
	// 	),
	// 	mcp.WithBoolean("force",
	// 		mcp.Description("是否强制提交（忽略lint错误）")),
	// )
	// s.AddTool(tool, gitCommitAndPushHandler)
	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
func currentTimeHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	timezone, ok := request.Params.Arguments["timezone"].(string)
	if !ok {
		return mcp.NewToolResultText("timezone must be a string"), nil
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("parse timezone with error: %v", err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf(`current time is %s`, time.Now().In(loc))), nil
}
