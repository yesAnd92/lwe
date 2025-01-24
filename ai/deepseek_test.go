package ai

import "testing"

func TestDeepSeek_Chat(t *testing.T) {
	ds := &DeepSeek{}
	ds.Chat("hh", "") // 添加第二个参数为空字符串以满足函数签名要求
}
