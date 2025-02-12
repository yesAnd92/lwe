package gitcmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

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

// GetGitCommitMsg git commit msg from ai
func GetGitCommitMsg(dir string) string {

	//check and init agent
	agent := ai.NewAIAgent()

	//check git repo
	if !checkExistGitRepo(dir) {
		cobra.CheckErr(fmt.Sprintf("%s Not a git repo!", dir))
	}

	//get git diff
	fmt.Println("Get git diff...")
	diff := buildGitDiffReq(dir)

	if len(diff) == 0 {
		cobra.CheckErr("There is no changes!")
	}

	//send ai to summary
	fmt.Printf("AI is generating commit message...\n\n")
	resp, err := gitDiffSubmitToAi(diff, agent)
	if err != nil {
		cobra.CheckErr(err)
	}

	return buildCommitMsg(resp)

}

func CommitAndPush(dir, cmsg string) {

	fmt.Println("AI suggested commit message:")
	printCommitMsg(dir, cmsg)

	// git add and git commit
	addAndCommit(dir, cmsg)

	//push origin repo
	pushCommitOriginRepo(dir)
}

func addAndCommit(dir string, cmsg string) {
	//accept cmsg
	var accept bool
	promptConfirm := &survey.Confirm{
		Message: "Accept this commit?",
	}
	err := survey.AskOne(promptConfirm, &accept)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("Commit msg err: %v", err))
	}

	if accept {
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
		// no, exit
		os.Exit(0)
	}
	//highlight hint
	fmt.Println(text.Colors{text.FgGreen, text.Bold}.Sprint("\nSuccess commit!\n"))

}

func getAllChangedFiles(dir string) string {
	var cmdline = fmt.Sprintf(STATUS_TPL_SHORT, dir)

	result := utils.RunCmd(cmdline, time.Second*5)
	if result.Err() != nil {
		cobra.CheckErr(result.Err().Error())
	}
	return result.String()
}

func printCommitMsg(dir, msg string) {

	files := getAllChangedFiles(dir)

	t := table.NewWriter()
	// Define the header row and set the style of the header cells, here the header color is set to blue
	headerRow := table.Row{"Files", "Commit msg"}
	for i := range headerRow {
		headerRow[i] = text.Colors{text.FgGreen}.Sprint(headerRow[i])
	}
	t.AppendHeader(headerRow)
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{files, msg})
	t.Render()
}

func pushCommitOriginRepo(dir string) {

	//accept cmsg
	var accept bool
	promptConfirm := &survey.Confirm{
		Message: "Accept this commit and push to origin repo?",
	}
	err := survey.AskOne(promptConfirm, &accept)
	if err != nil {
		cobra.CheckErr(fmt.Sprintf("Confirm Commit msg err: %v", err))
	}

	if accept {
		// yes, push to origin repo
		gitPushCmd := fmt.Sprintf(GIT_PUSH, dir)
		addResult := utils.RunCmd(gitPushCmd, time.Second*30)
		if addResult.Err() != nil {
			fmt.Print(addResult.String())
			cobra.CheckErr(addResult.Err())
		}
		//output push result
		fmt.Printf("\n%s\n", addResult.String())
		fmt.Println(text.Colors{text.FgGreen, text.Bold}.Sprint("\nSuccess push origin Repo!\n"))
	} else {
		// no, exit
		os.Exit(0)
	}
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
