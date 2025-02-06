package ai

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/config"
	"github.com/yesAnd92/lwe/utils"
)

type AIChat interface {
	Chat(ctx, prompt string) (string, error)
}

type AIAgent struct {
	AiChat AIChat
}

// Chat Wrapper AIChat's Chat method ,enhance it
func (aiAgent *AIAgent) Chat(ctx, prompt string) (string, error) {

	//Although it is clearly required in the prompt to return a clean JSON format
	//some APIs include the ```json``` tag when returning results.
	resp, err := aiAgent.AiChat.Chat(ctx, prompt)
	if err == nil {
		resp = utils.RemoveJSONTags(resp)
	}

	return resp, err
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
