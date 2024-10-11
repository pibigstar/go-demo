package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"testing"

	"github.com/sashabaranov/go-openai"
)

// input ：$0.0015 / 1K tokens
// output：$0.002 / 1K tokens
// 1K token 大约 500个汉字 或 750个英文单词
// token计算：https://platform.openai.com/tokenizer
// 支持的token最大长度为4097

const token = "xxx"

func TestChatCompletion(t *testing.T) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "请对以下电商工单数据进行分析总结，回答必须以该 {\"tag\": \"\"} json格式输出, tag为问题详细的分析总结结果， 字数必须控制在10到20内，不要给出额外的解释说明\n工单数据：看不到订单。说系统异常。只能看到物流说我已经签收。我找不到订单。咋截图啊。我拍的是车载充气泵。物流显示我签收了",
				},
			},
			MaxTokens: 200, // 限制返回的最大token数
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	// {"tag": "系统异常，订单未找到"}
	fmt.Println(resp.Choices[0].Message.Content)

	// {chatcmpl-7ssDS43fJQaEkeQE1JW5yZrmJunwm chat.completion 1693312474 gpt-3.5-turbo-0613 [{0 {assistant {"tag": "系统异常，订单未找到"}  <nil>} stop}] {138 12 150}}
	fmt.Println(resp)
}

func TestChatCompletionWithContext(t *testing.T) {
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "请对以下电商工单数据进行分析总结，回答必须以该 {\"tag\": \"\"} json格式输出, tag为问题详细的分析总结结果， 字数必须控制在10到20内，不要给出额外的解释说明",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	answer := resp.Choices[0].Message.Content
	fmt.Println(answer)

	resp, err = client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: answer, // 将上一次的回答也带过去
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "工单数据：看不到订单。说系统异常。只能看到物流说我已经签收。我找不到订单。咋截图啊。我拍的是车载充气泵。物流显示我签收了",
				},
			},
		},
	)
}

// 流式传输，NLG
func TestChatStreamingCompletion(t *testing.T) {
	c := openai.NewClient(token)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 500,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello ",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

// 测试向量计算与相似度识别
func TestEmbedding(t *testing.T) {
	client := openai.NewClient(token)

	input := "优惠券用户在个人主页没有找到"
	input2 := "用户反馈优惠券未收到"
	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: []string{input, input2},
			Model: openai.AdaEmbeddingV2,
		})

	if err != nil {
		fmt.Printf("CreateEmbeddings error: %v\n", err)
		return
	}

	fmt.Printf("%+v \n", resp)

	vectors := resp.Data[0].Embedding // []float32 with 1536 dimensions

	fmt.Printf("长度：%d \n", len(vectors))
	// [-0.007011389 -0.020194992 -0.0096179675 0.013995376 -0.0136117535 0.008035525 -0.0020362828 -0.021934995 -0.013426793 -0.028251069] ... [-0.0035279584 0.008453399 -0.007994422 -0.01490648 -0.016687585 0.023812005 -0.006165364 -0.026154844 0.008432848 0.010522221]
	fmt.Println(vectors[:10], "...", vectors[len(vectors)-10:])

	fmt.Println("相似度：", CosineSimilarity(resp.Data[0].Embedding, resp.Data[1].Embedding))
}

// CosineSimilarity 向量空间余弦相似度
func CosineSimilarity(a []float32, b []float32) float64 {
	var (
		aLen  = len(a)
		bLen  = len(b)
		s     = 0.0
		sa    = 0.0
		sb    = 0.0
		count = 0
	)
	if aLen > bLen {
		count = aLen
	} else {
		count = bLen
	}
	for i := 0; i < count; i++ {
		if i >= bLen {
			sa += math.Pow(float64(a[i]), 2)
			continue
		}
		if i >= aLen {
			sb += math.Pow(float64(b[i]), 2)
			continue
		}
		s += float64(a[i] * b[i])
		sa += math.Pow(float64(a[i]), 2)
		sb += math.Pow(float64(b[i]), 2)
	}
	return s / (math.Sqrt(sa) * math.Sqrt(sb))
}
