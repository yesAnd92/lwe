package sync

import (
	"testing"
)

func Test_compareDir(t *testing.T) {
	type args struct {
		sourceDir string
		targetDir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				sourceDir: "/Users/wangyj/ideaProject/my/lwe",
				targetDir: "/Users/wangyj/Desktop/lwe_copy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compareDir(tt.args.sourceDir, tt.args.targetDir)
		})
	}
}
