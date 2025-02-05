package gitcmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/ai"
	"github.com/yesAnd92/lwe/ai/prompt"
	"github.com/yesAnd92/lwe/utils"
)

type CommitMsg struct {
	Type        string `json:"type"`        // Type of the commit, e.g., feat, fix, docs, etc.
	Scope       string `json:"scope"`       // Scope of the impact (optional)
	Description string `json:"description"` // Brief description of the commit
}

type CommitData struct {
	CommitMsg      []CommitMsg `json:"commitMsg"` // List of commit messages
	OptionalBody   string      `json:"optionalBody"`
	OptionalFooter string      `json:"optionalFooter"`
}

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

	var commitData CommitData

	err = json.Unmarshal([]byte(resp), &commitData)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("parse response error:%v", err))
	}

}

func buildGitDiffReq() string {
	//git diff result

	var cmdline = GIT_DIFF
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
