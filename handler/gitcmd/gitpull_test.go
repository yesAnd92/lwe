package gitcmd

import (
	"fmt"
	"testing"
)

func TestCheckRepoClean(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{dir: "."},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clean, msg := checkRepoClean(tt.args.dir)
			if !clean {
				fmt.Println(msg)
			}
		})
	}
}

func TestUpdateAllGitRepo(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{args: args{dir: "/Users/wangyj/ideaProject/my/lwe"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateAllGitRepo(tt.args.dir)
		})
	}
}
