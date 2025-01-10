package gitcmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/ai"
	"github.com/yesAnd92/lwe/ai/prompt"
	"github.com/yesAnd92/lwe/utils"
	"strings"
)

type GitSummaryPromptResp struct {
	RepoSummary []RepoSummary `json:"repo_summary"`
}

type RepoSummary struct {
	Repo      string   `json:"repo"`
	Summary   []string `json:"summary"`
	SummaryCN []string `json:"summary_cn"`
}

// GitLogSummary summary commit log
func GitLogSummary(detail bool, dir, committer, start, end string) {

	//check and init agent
	agent := ai.NewAIAgent()

	sb := buildGitLogReq(detail, dir, committer, start, end)

	//send ai to summary
	resp, err := logSubmitToAi(sb, agent)
	if err != nil {
		cobra.CheckErr(err)
	}

	promptResp := parseResp(resp)

	consoleResult(promptResp)

}

func consoleResult(promptResp *GitSummaryPromptResp) {
	//输出
	for i, repoSum := range promptResp.RepoSummary {
		fmt.Printf("#%d. %s", i+1, repoSum.Repo)

		fmt.Print("\nEN:\n")
		for no, s := range repoSum.Summary {
			fmt.Printf("%d. %s\n", no, s)
		}

		fmt.Print("\nCN:\n")

		for no, s := range repoSum.SummaryCN {
			fmt.Printf("%d. %s\n", no, s)
		}

		//each repo split by to blank line
		fmt.Print("\n\n")
	}
}

func parseResp(resp string) *GitSummaryPromptResp {
	promptResp := &GitSummaryPromptResp{}

	err := json.Unmarshal([]byte(resp), promptResp)
	if err != nil {
		cobra.CheckErr(err)
	}
	return promptResp
}

func buildGitLogReq(detail bool, dir string, committer string, start string, end string) string {
	var res []string

	//相对路径转换成绝对路径进行处理
	dir = utils.ToAbsPath(dir)

	//递归找到所有的git仓库
	findGitRepo(dir, &res)

	//recentN not more than 1000
	var recentN int16 = 1000

	//get all branch commit log
	var allBranch = true

	//text to be submitted to AI for summarization
	var sb strings.Builder

	//遍历获取每个仓库的提交信息
	for _, gitDir := range res {
		commitLogs, err := GetAllGitRepoCommitLog(detail, recentN, gitDir, committer, start, end, allBranch)
		if err != nil {
			cobra.CheckErr(err)
		}
		if len(*commitLogs) == 0 {
			break
		}
		for _, repoLogs := range *commitLogs {
			sb.WriteString(repoLogs.RepoName + "\n")
			for _, log := range *repoLogs.CommitLogs {
				logMsg := strings.TrimSpace(log.CommitMsg)
				if len(logMsg) == 0 {
					continue
				}
				sb.WriteString(logMsg + "\n")
			}
			sb.WriteString("\n")
		}

	}
	return sb.String()
}

func logSubmitToAi(ctx string, aiAgent *ai.AIAgent) (string, error) {

	content := prompt.LogSummaryPrompt + "\n" + ctx
	//submit to the AI using the preset prompt
	resp, err := aiAgent.AiChat.Chat(content)
	return resp, err
}
