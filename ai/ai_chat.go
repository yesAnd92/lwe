package ai

type AIChat interface {
	Chat(ctx string) (string, error)
}

type AIAgent struct {
	AiChat AIChat
}

func NewAIAgent() *AIAgent {

	// TODO: 2025/1/7 Read the configuration file and decide which AI to use.

	agent := AIAgent{AiChat: &DeepSeek{}}
	return &agent
}
