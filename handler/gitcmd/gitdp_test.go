package gitcmd

import (
	"testing"
)

func TestGetDeployCommand(t *testing.T) {
	type args struct {
		updatePj string
		msg      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1",
			args: args{
				updatePj: "new-media-rft",
				msg:      "",
			}},
		{name: "case2",
			args: args{
				updatePj: "new-media-rft,new-media-mp",
				msg:      "",
			}},
		{name: "bad-case-",
			args: args{
				updatePj: "new-media-rft, ,",
				msg:      "",
			}},
		{name: "bad-case-中文逗号",
			args: args{
				updatePj: "new-media-rft，new-media-mp,,,,",
				msg:      "",
			}},
		{name: "bad-case-连续多逗号兼容,@符号",
			args: args{
				updatePj: "@new-media-rft,new-media-mp,new-media-mpclient,,,",
				msg:      "",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagInfo := BuildTagInfo(tt.args.updatePj, tt.args.msg)
			PrintTplResult(tagInfo)
		})
	}
}
