package gitcmd

import (
	"fmt"
	"github.com/yesAnd92/lwe/ai"
	"os"
	"testing"
)

func TestGitCommitMsg(t *testing.T) {
	msg := GetGitCommitMsg(".")
	fmt.Println(msg)

}

func Test_gitDiffSubmitToAi(t *testing.T) {
	type args struct {
		diffFile string
		aiAgent  *ai.AIAgent
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "demo",
			args: args{diffFile: "../../testdata/diff.log",
				aiAgent: ai.NewAIAgent()},
		},
	}
	for _, tt := range tests {
		content, err := os.ReadFile(tt.args.diffFile)
		if err != nil {
			panic(err)
		}
		ctx := string(content)
		fmt.Println(ctx)
		got, err := gitDiffSubmitToAi(ctx, tt.args.aiAgent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(got)
	}
}

func Test_buildCommitMsg(t *testing.T) {
	type args struct {
		resp string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "",
			args: args{resp: `{
						  "commitMsg": [
							{
							  "type": "feat",
							  "scope": "ai",
							  "description": "Add Siliconflow as a new AI agent type"
							},
							{
							  "type": "refactor",
							  "scope": "deepseek",
							  "description": "Rename Send function to dsSend and update response struct to CommonResponse"
							},
							{
							  "type": "test",
							  "scope": "deepseek",
							  "description": "Enhance DeepSeek Chat test with config initialization and proper error handling"
							}
						  ]
				}`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildCommitMsg(tt.args.resp)
			fmt.Println(got)
		})
	}
}

func Test_printCommitMsg(t *testing.T) {
	// 测试用例 1
	msg1 := "Commit message 1"
	printCommitMsg(".", msg1)

	// 测试用例 2
	msg2 := "Another commit message"
	printCommitMsg(".", msg2)
}
