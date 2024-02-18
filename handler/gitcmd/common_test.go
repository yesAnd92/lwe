package gitcmd

import (
	"fmt"
	"testing"
)

func TestListRepoAllBranch(t *testing.T) {
	type args struct {
		repo string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				repo: ".",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListRepoAllBranch(tt.args.repo)
			fmt.Println(got)
		})
	}
}
