package dp

import (
	"fmt"
	"net/http"
	"testing"
)

// 构建部署信息
func TestBuildDdpInfo(t *testing.T) {
	type args struct {
		updatePj string
		msg      string
		dir      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case1",
			args: args{
				updatePj: "new-media-rft",
				dir:      ".",
				msg:      "升级广播电视",
			}},
		{name: "case2",
			args: args{
				updatePj: "new-media-rft,new-media-mp",
				dir:      ".",
				msg:      "",
			}},
		{name: "bad-case-",
			args: args{
				updatePj: "new-media-rft, ,",
				dir:      ".",
				msg:      "",
			}},
		{name: "bad-case-中文逗号",
			args: args{
				updatePj: "new-media-rft，new-media-mp,,,,",
				dir:      ".",
				msg:      "",
			}},
		{name: "bad-case-连续多逗号兼容,@符号",
			args: args{
				updatePj: "@new-media-rft,new-media-mp,new-media-mpclient,,,",
				dir:      ".",
				msg:      "",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagInfo := BuildDdpInfo(tt.args.updatePj, tt.args.dir, tt.args.msg)
			PrintTplResult(tagInfo)
		})
	}
}

// 提交工单
func TestSubmitWorkSheet(t *testing.T) {
	dpObj := &DpObj{
		Msg:         "测试工单，请勿理会！，会自行撤单2",
		Tag:         "tag-rft-20230709-1234",
		PjName:      "冀云",
		PjNO:        "3",
		ModuleNames: "new-media-rft,new-media-mp",
		ModuleDesc:  "广播电视,媒体号",
		Username:    "wangyingjie@pdmi.cn",
		Pwd:         "Opspaas1234",
	}
	cookieJar := LoginItsm(dpObj)
	sn := SubmitWorkSheet(cookieJar, dpObj)
	fmt.Printf("工单号：%s", sn)
}

// 登录
func TestLoginItam(t *testing.T) {
	dpObj := &DpObj{
		Username: "wangyingjie@pdmi.cn",
		Pwd:      "Opspaas1234",
	}
	LoginItsm(dpObj)
}

func Test_gitRemoteRepo(t *testing.T) {
	repo, repoNo := getGitRemoteRepo(".")
	fmt.Printf("项目名；%s;项目编号：%s", repo, repoNo)
}

// 构建生成提交工单所需参数
func Test_buildSubmitWorksheetParam(t *testing.T) {
	type args struct {
		obj *DpObj
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{obj: &DpObj{
			Msg:         "哈哈哈哈",
			Tag:         "tag-rft-20230709-1234",
			PjName:      "冀云",
			PjNO:        "9527",
			ModuleNames: "new-media-rft,new-media-mp",
			ModuleDesc:  "广播电视,媒体号",
			Username:    "yesand",
			Pwd:         "yesand",
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildSubmitWorksheetParam(tt.args.obj); got != tt.want {
				fmt.Println(got)
			}
		})
	}

}

// 解析工单号
func Test_parseWorkSheetNo(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{body: []byte(`{
    "result": true,
    "code": "OK",
    "message": "success",
    "data": {
        "sn": "REQ20230706000002",
        "id": 463
    }
}`)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseWorkSheetNo(tt.args.body); got != tt.want {
				fmt.Println(got)
			}
		})
	}
}

func Test_getProject11(t *testing.T) {
	dpObj := &DpObj{
		Username: "wangyingjie@pdmi.cn",
		Pwd:      "Opspaas1234",
	}
	type args struct {
		client http.CookieJar
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{args: args{client: LoginItsm(dpObj)},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProject11(tt.args.client); got != tt.want {
				t.Errorf("getProject11() = %v, want %v", got, tt.want)
			}
		})
	}
}
