package gitcmd

import (
	"fmt"
	"github.com/yesAnd92/lwe/ai"
	"os"
	"testing"
)

func TestGitCommitMsg(t *testing.T) {
	GitCommitMsg()
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
