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
	SysPrompt = `
你现在需要处理电商产品描述生成任务，具体要求如下：

### 输入的信息
1. desc_all: 文字数据：请对其进行预处理，过滤重复文本、无实际意义的虚词（如“啊”“哦”“嗯”）、特殊符号（如#*&￥）及冗余表述，保留核心产品特征、功能或优势相关内容。
2. base_name:明确核心产品名称（如蓝牙耳机、充电宝、无人机），是内容聚焦的核心依据。

### 角色与原则
你是一位熟悉消费者心理、从业5年以上的资深电商销售，擅长提炼产品核心价值。描述需**真实克制**，拒绝无意义夸大，同时突出吸引力——让普通产品显得更值得购买。
输入尽量不涉及具体参数, 类似, 多大, 多高, 多快

### 输出要求
需结合预处理后的文字数据与basename，生成以下内容：
1. title：≤18字。需清晰点明产品，突出核心优点，语言简洁有力、易被注意（避免平淡表述）。
2. desc：≤50字。聚焦产品实用价值、性价比或独特优势，体现“值钱感”“物超所值”，语言通俗易懂。

### 格式约束
输出必须为严格JSON格式, 字段不可缺失, 不可新增, 示例如下：
{
    "title": "替换为生成的标题",
    "desc": "替换为生成的描述"
}
`
)

func getClient() *openai.Client {
	ak := os.Getenv("DP_AK")
	url := os.Getenv("DP_BASE_URL")
	fmt.Printf("output: %v, %v \n", ak, url)

	config := openai.DefaultConfig(ak)
	config.BaseURL = url
	return openai.NewClientWithConfig(config)
}

func One(ctx context.Context, context string) string {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: SysPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: context,
		},
	}

	resp, err := getClient().CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    ModelDPV3,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	//fmt.Println(resp)
	return resp.Choices[0].Message.Content
}

// 需要保留并累加完整的对话历史，每次交互时将新的用户提问和 AI 回复依次追加到消息列表中，确保上下文连贯性。
func More() {
	client := getClient()
	// 初始化对话历史（老对话的初始状态）
	history := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "如何使用Go语言读取JSON文件？", // 历史提问
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "可以使用encoding/json包的Unmarshal函数...", // 历史回复
		},
	}

	// 用户新提问：追加到对话历史
	history = append(history, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "那如何写入JSON文件呢？", // 新提问
	})

	// 发起持续对话请求（传入完整历史）
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: history, // 传入累计的对话历史
		},
	)

	// 将AI的新回复追加到历史，用于下一次对话
	// todo 中间需要有中断, 进行输入才行
	if err == nil && len(resp.Choices) > 0 {
		history = append(history, resp.Choices[0].Message)
	}
}
