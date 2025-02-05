package ai

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/config"
)

type AIChat interface {
	Chat(ctx, prompt string) (string, error)
}

type AIAgent struct {
	AiChat AIChat
}

type AIName string

const (
	Deepseek    = "deepseek"
	Siliconflow = "siliconflow"
)

func NewAIAgent() *AIAgent {

	// Read the configuration file and decide which AI to use.
	config := config.InitConfig()
	ai := config.Ai

	var agent AIAgent

	switch ai.Name {
	case Deepseek:
		agent = AIAgent{AiChat: &DeepSeek{}}
		break
	case Siliconflow:
		agent = AIAgent{AiChat: &SiliconFlow{}}
		break
	default:
		cobra.CheckErr("AI configuration is missing or incorrect.")
	}

	return &agent
}
