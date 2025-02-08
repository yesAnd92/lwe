package gitcmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strings"
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

// GitCommitMsg git commit msg from ai
func GitCommitMsg(dir string) string {

	//check and init agent
	agent := ai.NewAIAgent()

	diff := buildGitDiffReq(dir)

	if len(diff) == 0 {
		cobra.CheckErr("There is no changes!")
	}

	//send ai to summary
	resp, err := gitDiffSubmitToAi(diff, agent)
	if err != nil {
		cobra.CheckErr(err)
	}

	return buildCommitMsg(resp)

}

func CommitAndPush(dir, cmsg string) {

	fmt.Println("AI suggested commit msg:")
	fmt.Println(cmsg)
	//accept cmsg
	var accept bool
	promptConfirm := &survey.Confirm{
		Message: fmt.Sprintf("Accept this commit?"),
	}
	err := survey.AskOne(promptConfirm, &accept)
	if err != nil {
		fmt.Println("Commit msg err:", err)
		return
	}

	if accept {
		// yes git add and git commit

		addCmd := fmt.Sprintf(GIT_ADD, dir)
		addResult := utils.RunCmd(addCmd, time.Second*30)
		if addResult.Err() != nil {
			cobra.CheckErr(addResult.Err())
		}

		cmsgCmd := fmt.Sprintf(GIT_COMMIT, dir, cmsg)
		gcmsgReulst := utils.RunCmd(cmsgCmd, time.Second*30)
		if gcmsgReulst.Err() != nil {
			cobra.CheckErr(gcmsgReulst.Err())
		}
	} else {
		//结束
		return
	}
}

func pushCommitOriginRepo(cmsg string) {

	fmt.Println("")
	//accept cmsg
	var accept bool
	promptConfirm := &survey.Confirm{
		Message: fmt.Sprintf("Accept this commit and push to origin repo?"),
	}
	err := survey.AskOne(promptConfirm, &accept)
	if err != nil {
		fmt.Println("confirm commit msg err:", err)
		return
	}

	if accept {
		// yes
		// TODO: 2025/2/6 push to origin repo
	}
	fmt.Println(accept)
}

func buildCommitMsg(resp string) string {
	var commitData CommitData

	err := json.Unmarshal([]byte(resp), &commitData)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("parse %s \n error:%v", resp, err))
	}

	var cmsg strings.Builder

	for _, msg := range commitData.CommitMsg {
		line := fmt.Sprintf("%s(%s): %s\n", msg.Type, msg.Scope, msg.Description)
		cmsg.WriteString(line)
	}

	if len(commitData.OptionalBody) > 0 {
		cmsg.WriteString("\n")
		cmsg.WriteString(commitData.OptionalBody)
	}

	if len(commitData.OptionalFooter) > 0 {
		cmsg.WriteString("\n")
		cmsg.WriteString(commitData.OptionalFooter)
	}
	return cmsg.String()
}

func buildGitDiffReq(dir string) string {
	//git diff result

	var cmdline = fmt.Sprintf(GIT_DIFF, dir)
	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		cobra.CheckErr(result.Err())
	}
	return result.String()
}

func gitDiffSubmitToAi(ctx string, aiAgent *ai.AIAgent) (string, error) {

	//submit to the AI using the preset prompt
	resp, err := aiAgent.Chat(ctx, prompt.GitDiffPrompt)
	return resp, err
}
