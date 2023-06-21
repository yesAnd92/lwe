package gitcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lwe/utils"
	"os"
	"strings"
	"text/template"
	"time"
)

const (

	//gdp
	DP_TAG_TPL  = "tag-%s-%s"
	DP_TAG      = "git tag %s"
	DP_PUSH_TAG = "git push origin %s"
)

var resultTpl = `
工单：
部署升级命令： @{{.PjNames}}
tag:{{.Tag}}
升级项目：
升级哪些功能：{{.Msg}}
升级时长：5分钟
影响模块：{{.AffectPj}}
是否经过开发自测：是
是否经过测试验证： 否
是否对服务器有影响：否
`

type GdpObj struct {
	Msg      string
	Tag      string
	PjNames  string
	AffectPj string
}

func BuildTagInfo(updatePj, msg string) *GdpObj {

	//兼容空格，中文逗号，去掉@
	updatePj = strings.NewReplacer(" ", "", "，", ",", "@", "").Replace(updatePj)

	projectArr := strings.Split(updatePj, ",")

	pjNamesBdArr := []string{}
	affectPjBdArr := []string{}

	pjNameArr := []string{}
	for _, pj := range projectArr {
		//兼容空内容
		if len(pj) == 0 {
			continue
		}
		shortName, ok := checkProjectName(pj)
		if !ok {
			cobra.CheckErr(fmt.Sprintf("please check [%s]", pj))
		}

		split := strings.Split(shortName, ";")

		//多个只取前3个作为摘要
		if len(pjNameArr) < 3 {
			pjNameArr = append(pjNameArr, split[0])
		}
		pjNamesBdArr = append(pjNamesBdArr, pj)
		affectPjBdArr = append(affectPjBdArr, split[1])
	}

	tagDate := time.Now().Format("20060102-1504")

	tag := fmt.Sprintf(DP_TAG_TPL, strings.Join(pjNameArr, ""), tagDate)

	// TODO: 2023/6/21 正则校验下结果

	return &GdpObj{
		Msg:      msg,
		Tag:      tag,
		PjNames:  strings.Join(pjNamesBdArr, ","),
		AffectPj: strings.Join(affectPjBdArr, ","),
	}
}

func CommitAndPushTag(tagInfo, path string) {

	if len(path) > 0 {
		//指定了目录，切换到指定目录执行命令
		if err := os.Chdir(path); err != nil {
			cobra.CheckErr(err)
		}
	}

	var tagCmdline = fmt.Sprintf(DP_TAG, tagInfo)

	//create tag
	result := utils.RunCmd(tagCmdline, time.Second*30)
	if result.Err() != nil {
		cobra.CheckErr(result.Err().Error())
	}
	fmt.Printf("%s\n", tagCmdline)

	var tagPushCmdline = fmt.Sprintf(DP_PUSH_TAG, tagInfo)

	//push tag
	fmt.Printf("%s...\n", tagPushCmdline)
	pushRe := utils.RunCmd(tagPushCmdline, time.Second*30)
	if pushRe.Err() != nil {
		cobra.CheckErr(pushRe.Err().Error())
	}
	fmt.Println(pushRe.String())

}

func PrintTplResult(obj *GdpObj) {
	tpl := template.Must(template.New("gdpTpl").Funcs(utils.TemplateFunc).Parse(resultTpl))
	tpl.Execute(os.Stdout, obj)
}

func checkProjectName(pj string) (string, bool) {

	if s, ok := pjNameMap[pj]; ok {
		return s, ok
	}
	return "", false
}
