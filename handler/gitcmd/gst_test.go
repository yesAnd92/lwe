package gitcmd

import "testing"

func TestGetAllGitRepoStatus(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{dir: "."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllGitRepoStatus(tt.args.dir)
		})
	}
}
