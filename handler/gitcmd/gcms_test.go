package gitcmd

import (
	"fmt"
	"github.com/yesAnd92/lwe/ai"
	"testing"
)

func TestGitCommitMsg(t *testing.T) {
	GitCommitMsg()
}

func Test_gitDiffSubmitToAi(t *testing.T) {
	type args struct {
		ctx     string
		aiAgent *ai.AIAgent
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "demo",
			args: args{ctx: "",
				aiAgent: ai.NewAIAgent()},
		},
	}
	for _, tt := range tests {
		got, err := gitDiffSubmitToAi(tt.args.ctx, tt.args.aiAgent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(got)
	}
}
