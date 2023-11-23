package pdf

import "testing"

func Test_parseMergeArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				args: []string{"", "bookletTest.pdf", "github.png", "testdata/*.jpg"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re, err := parseMergeArg(tt.args.args)
			if err != nil {
				t.Fail()
			}
			if len(re) != 3 {
				t.Fail()
			}
		})
	}
}
