package ai

type AIChat interface {
	Chat(ctx string) (string, error)
}

type AIProxy struct {
	AiChat AIChat
}

func NewAIProxy() *AIProxy {

	// TODO: 2025/1/7 Read the configuration file and decide which AI to use.

	proxy := AIProxy{AiChat: &DeepSeek{}}
	return &proxy
}
