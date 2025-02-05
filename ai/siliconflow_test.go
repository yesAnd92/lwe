package ai

import (
	"fmt"
	"github.com/yesAnd92/lwe/config"
	"testing"
)

func TestSiliconFlow_Chat(t *testing.T) {
	config.InitConfig()
	sf := &SiliconFlow{}
	chat, err := sf.Chat("hello,siliconflow", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(chat)
}
