package gitcmd

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/ai"
	"github.com/yesAnd92/lwe/ai/prompt"
	"github.com/yesAnd92/lwe/utils"
)

// GitCommitMsg git commit msg from ai and pull request to github
func GitCommitMsg() {

	//check and init agent
	agent := ai.NewAIAgent()

	sb := buildGitDiffReq()

	//send ai to summary
	resp, err := gitDiffSubmitToAi(sb, agent)
	if err != nil {
		cobra.CheckErr(err)
	}

	fmt.Print(resp)

}

func buildGitDiffReq() string {
	//获取git diff

	var cmdline = fmt.Sprintf("%s %s %s", "git", "diff", "HEAD")
	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		cobra.CheckErr(result.Err())
	}
	return result.String()
}

func gitDiffSubmitToAi(ctx string, aiAgent *ai.AIAgent) (string, error) {

	//submit to the AI using the preset prompt
	resp, err := aiAgent.AiChat.Chat(ctx, prompt.GitDiffPrompt)
	return resp, err
}
