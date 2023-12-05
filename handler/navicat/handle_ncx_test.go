package navicat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseNcx(t *testing.T) {

	tests := []struct {
		name   string
		data   []byte
		expect string
	}{
		// test cases
		{
			name:   "pwd case",
			data:   []byte(`<Connections Ver="1.5"><Connection ConnectionName="pwd case" ProjectUUID="" ConnType="MYSQL"  Host="127.0.0.1" Port="3306" UserName="root" Password="B75D320B6211468D63EB3B67C9E85933"/></Connections>`),
			expect: `{"Conns":[{"ConnectionName":"pwd case","ConnType":"MYSQL","Host":"127.0.0.1","UserName":"root","Port":"3306","Password":"This is a test"}],"Version":"1.5"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCons, err := ParseNcx(tt.data)
			result, err := json.Marshal(gotCons)
			if err != nil {
				t.Errorf("ParseNcx() error %e", err)
				return
			}
			if !reflect.DeepEqual(string(result), tt.expect) {
				t.Errorf("except:%s\n	actual:%s", tt.expect, result)
			}
		})
	}
}

func Test_decryptPwd(t *testing.T) {
	type args struct {
		encryptTxt string
		expect     string
	}
	tests := []struct {
		name string
		args args
	}{
		// test cases.
		{
			name: "case 1",
			args: args{
				encryptTxt: "833E4ABBC56C89041A9070F043641E3B",
				expect:     "123456",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decryptPwd(tt.args.encryptTxt)
			if !reflect.DeepEqual(got, tt.args.expect) {
				t.Errorf("except:%s\n actual:%s", tt.args.expect, got)
			}
		})
	}
}
