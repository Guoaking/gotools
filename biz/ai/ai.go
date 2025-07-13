package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

/**
@description
@date: 07/13 19:28
@author Gk
**/

const (
	ModelDPV3 = "deepseek-chat"
	ModelDPR1 = "deepseek-reasoner "
)

func o1() {

	ak := os.Getenv("DP_AK")
	url := os.Getenv("DP_BASE_URL")
	fmt.Printf("output: %v, %v \n", ak, url)

	config := openai.DefaultConfig(ak)
	config.BaseURL = url
	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: ModelDPV3,
			Messages: []openai.ChatCompletionMessage{
				{
					// 系统消息角色，用于设置对话的上下文和行为。通常由开发者提供，用于指导模型的行为，例如设定模型的身份、风格、任务等。
					//使用场景：可以用来定义模型的基本属性，如要求模型以某种特定的语气回复，或者告知模型当前的任务是什么。
					// 回复格式: {"
					//"}
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					//
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// todo 生成 json ?保存?

	fmt.Println(resp.Choices[0].Message.Content)
}
