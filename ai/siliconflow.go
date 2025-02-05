package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/config"
	"net/http"
)

type SiliconFlow struct {
}

func (ds *SiliconFlow) Chat(ctx, prompt string) (string, error) {
	resp := sfSend(ctx, prompt)
	return resp, nil
}

func sfSend(ctx, prompt string) string {

	aiConfig := config.GlobalConfig.Ai
	url := aiConfig.BaseUrl
	apiKey := aiConfig.ApiKey
	model := aiConfig.Model

	method := "POST"

	// to build message for request
	message := []map[string]interface{}{
		{
			"content": ctx,
			"role":    "user",
		},
		{
			"content": prompt,
			"role":    "system",
		},
	}
	requestData := map[string]interface{}{
		"messages":          message,
		"model":             model,
		"frequency_penalty": 0.5,
		"max_tokens":        2048,
		"response_format": map[string]interface{}{
			"type": "text",
		},
		"stop":   nil,
		"stream": false,
	}

	// 将map转换为JSON格式的字节切片
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		cobra.CheckErr(err)
	}

	// 创建一个io.Reader用于请求体
	payload := bytes.NewReader(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		cobra.CheckErr(err)
	}

	auth := fmt.Sprintf("Bearer %s", apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", auth)

	resp, err := client.Do(req)
	if err != nil {
		cobra.CheckErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		cobra.CheckErr(fmt.Sprintf("request fail ,statusCode: %d", resp.StatusCode))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		cobra.CheckErr(fmt.Sprintf("parse fail: %v", err))
	}

	if errorData, ok := result["error"]; ok {
		cobra.CheckErr(fmt.Sprintf("request error: %v", errorData))
	}

	dsResp := &CommonResponse{}
	if err := mapstructure.Decode(result, &dsResp); err != nil {
		cobra.CheckErr(err)
	}
	return dsResp.Choices[0].Message.Content
}
