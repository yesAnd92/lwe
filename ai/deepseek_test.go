package ai

import (
	"fmt"
	"github.com/yesAnd92/lwe/config"
	"testing"
)

func TestDeepSeek_Chat(t *testing.T) {
	config.InitConfig()
	ds := &DeepSeek{}
	chat, err := ds.Chat("hello,deepseek", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(chat)
}
