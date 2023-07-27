package gitcmd

import (
	"fmt"
	"testing"
)

var token = ""

func TestGetRepositories(t *testing.T) {
	urls := getRepoList("", token)
	for _, url := range urls {
		fmt.Println(url)
	}
}

func TestCloneGroup(t *testing.T) {
	type args struct {
		url       string
		token     string
		targetDir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1",
			args: args{
				url:       "",
				token:     token,
				targetDir: "",
			}},
		{name: "case2",
			args: args{
				url:       "",
				token:     token,
				targetDir: "",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CloneGroup(tt.args.url, tt.args.token, tt.args.targetDir)
		})
	}
}
