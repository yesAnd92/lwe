package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url    = "https://api.deepseek.com/chat/completions"
	apiKey = ""
)

type DeepSeek struct {
}

func (ds *DeepSeek) Chat(ctx string) (string, error) {
	resp := Send(ctx)
	return resp, nil
}

func Send(ctx string) string {

	method := "POST"

	// 使用map构建请求数据结构
	requestData := map[string]interface{}{
		"messages": []map[string]interface{}{
			{
				"content": ctx,
				"role":    "user",
			},
		},
		"model":             "deepseek-chat",
		"frequency_penalty": 0,
		"max_tokens":        2048,
		"presence_penalty":  0,
		"response_format": map[string]interface{}{
			"type": "text",
		},
		"stop":           nil,
		"stream":         false,
		"stream_options": nil,
		"temperature":    1,
		"top_p":          1,
		"tools":          nil,
		"tool_choice":    "none",
		"logprobs":       false,
		"top_logprobs":   nil,
	}

	// 将map转换为JSON格式的字节切片
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return ""
	}

	// 创建一个io.Reader用于请求体
	payload := bytes.NewReader(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	auth := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", auth)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	resp := &DeepSeekResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return resp.Choices[0].Message.Content
}
