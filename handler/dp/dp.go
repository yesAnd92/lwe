package dp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"lwe/utils"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"text/template"
	"time"
)

const (

	//dp
	DP_TAG_TPL     = "tag-%s-%s"
	DP_TAG         = "git tag %s"
	DP_PUSH_TAG    = "git push origin %s"
	DP_REMOTE_INFO = "git -C %s remote -v"

	//蓝鲸登录url
	BLUEKING_LOGIN_URL = "http://paas.hubpd.com/login/?c_url=/console/"
	//itsm登录
	ITSM_LOGIN_URL = "http://paas.hubpd.com/o/bk_itsm"
	//itsm工单提交
	ITSM_SUBMIT_WORKSHEET_URL = "http://paas.hubpd.com/o/bk_itsm/api/ticket/receipts/"
	//登录用的一个随机token
	BKLOGIN_CSRFTOKEN = "nTc7dR2PyKv7y92yffjtfLT4CFC3Laah"
)

type DpObj struct {
	WorkSheetSn string //工单号
	Msg         string //升级描述
	Tag         string //升级tag
	PjName      string //升级项目名
	PjNO        string //升级项目编号
	ModuleNames string //涉及模块名
	ModuleDesc  string //影响模块
	Username    string //用户名
	Pwd         string //密码
}

func BuildDdpInfo(updateModule, dir, msg string) *DpObj {
	//获取git repo信息
	pjName, pjNO := getGitRemoteRepo(dir)

	//兼容空格，中文逗号，去掉@
	updateModule = strings.NewReplacer(" ", "", "，", ",", "@", "").Replace(updateModule)

	ModuleArr := strings.Split(updateModule, ",")

	var moduleNamesBdArr []string
	var moduleDescArr []string

	var pjNameArr []string
	for _, pj := range ModuleArr {
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
		moduleNamesBdArr = append(moduleNamesBdArr, pj)
		moduleDescArr = append(moduleDescArr, split[1])
	}

	tagDate := time.Now().Format("20060102-1504")

	tag := fmt.Sprintf(DP_TAG_TPL, strings.Join(pjNameArr, ""), tagDate)

	// TODO: 2023/6/21 正则校验下结果

	return &DpObj{
		Msg:         msg,
		Tag:         tag,
		PjName:      pjName,
		PjNO:        pjNO,
		ModuleNames: strings.Join(moduleNamesBdArr, ","),
		ModuleDesc:  strings.Join(moduleDescArr, ","),
		Username:    username,
		Pwd:         pwd,
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

func PrintTplResult(obj *DpObj) {
	tpl := template.Must(template.New("dpTpl").Funcs(utils.TemplateFunc).Parse(dpResultTpl))
	tpl.Execute(os.Stdout, obj)
}

func checkProjectName(pj string) (string, bool) {

	if s, ok := moduleNameMap[pj]; ok {
		return s, ok
	}
	return "", false
}

func getGitRemoteRepo(path string) (string, string) {

	var remoteCmdline = fmt.Sprintf(DP_REMOTE_INFO, path)
	//create tag
	result := utils.RunCmd(remoteCmdline, time.Second*30)
	if result.Err() != nil {
		cobra.CheckErr("Please confirm the path is a git repository!")
	}
	resp := result.String()
	split := strings.Split(resp, "\n")

	//获取获取的项目名称
	pjName := ""
	if len(split) > 1 {
		pushSouce := split[1]
		pjName = pushSouce[strings.LastIndex(pushSouce, "/")+1 : strings.LastIndex(pushSouce, ".")]
	}
	if name, ok := pjNameMap[pjName]; ok {
		split := strings.Split(name, ";")
		return split[0], split[1]
	}
	cobra.CheckErr("This repository not in deploy list!")
	return "", ""
}

func buildSubmitWorksheetParam(obj *DpObj) string {
	tpl := template.Must(template.New("loginItamTpl").Funcs(utils.TemplateFunc).Parse(logInItamTpl))

	jsonParam := bytes.Buffer{}
	tpl.Execute(&jsonParam, obj)
	v := make(map[string]interface{})
	err1 := json.Unmarshal(jsonParam.Bytes(), &v)
	if err1 != nil {
		fmt.Println(err1)
	}
	re, err := json.Marshal(v)
	if err != nil {
		cobra.CheckErr("SubmitWorksheetParam error!")
	}
	return string(re)
}

func LoginItsm(obj *DpObj) http.CookieJar {

	// 创建一个 cookie jar,用于存放登录后返回的cookie，
	//后续调用接口使用
	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Timeout: 5 * time.Second,
		Jar:     jar,
	}

	param := fmt.Sprintf("username=%s&password=%s&csrfmiddlewaretoken=%s", obj.Username, obj.Pwd, BKLOGIN_CSRFTOKEN)
	req, _ := http.NewRequest("POST", BLUEKING_LOGIN_URL, strings.NewReader(param))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "bklogin_csrftoken="+BKLOGIN_CSRFTOKEN)
	resp, err := client.Do(req)
	if err != nil {
		cobra.CheckErr(err)
	}

	if resp.StatusCode != 200 {
		cobra.CheckErr("BlueKing login failed!")
	}
	//登录itsm
	itsmLoginResp, err := client.Get(ITSM_LOGIN_URL)
	if err != nil {
		cobra.CheckErr(err)
	}

	if itsmLoginResp.StatusCode != 200 {
		cobra.CheckErr("Itsm login failed!")
	}

	return client.Jar
}

// 仅作为测试cookie使用
func getProject11(cookiejar http.CookieJar) bool {
	upgradeUrl := "http://paas.hubpd.com/o/bk_itsm/api/service/projects/11/"
	upgradeReq, _ := http.NewRequest("GET", upgradeUrl, nil)
	client := http.Client{Jar: cookiejar}
	upgradeResp, err := client.Do(upgradeReq)
	if err != nil {
		fmt.Println(err)
	}
	defer upgradeResp.Body.Close()
	body, _ := ioutil.ReadAll(upgradeResp.Body)

	re := make(map[string]interface{})
	err = json.Unmarshal(body, &re)
	if err != nil {
		fmt.Println(err)
	}
	if result, ok := re["result"]; !ok || result != true {
		return false
	}
	return true
}

// SubmitWorkSheet 提交工单
func SubmitWorkSheet(cookieJar http.CookieJar, obj *DpObj) string {
	param := buildSubmitWorksheetParam(obj)
	submitReq, _ := http.NewRequest("POST", ITSM_SUBMIT_WORKSHEET_URL, strings.NewReader(param))
	cookies := cookieJar.Cookies(submitReq.URL)
	//提交工单除了cookie中的参数信息，header中还要增加csrfToken
	csrfToken := ""
	for _, cookie := range cookies {
		c := cookie.String()
		if strings.HasPrefix(c, "bkitsm_csrftoken=") {
			csrfToken = strings.TrimPrefix(c, "bkitsm_csrftoken=")
		}

	}
	fmt.Println(csrfToken)
	submitReq.Header.Add("Content-Type", "application/json;charset=UTF-8")
	submitReq.Header.Add("X-CSRFToken", csrfToken)
	client := http.Client{
		Jar: cookieJar,
	}
	upgradeResp, err := client.Do(submitReq)

	if err != nil {
		fmt.Println(err)
	}
	defer upgradeResp.Body.Close()
	body, _ := ioutil.ReadAll(upgradeResp.Body)

	sn := parseWorkSheetNo(body)

	//工单号放到obj中
	obj.WorkSheetSn = sn

	fmt.Println("Submit worksheet success...")
	return sn
}

func parseWorkSheetNo(body []byte) string {
	re := make(map[string]interface{})
	err := json.Unmarshal(body, &re)
	if err != nil {
		fmt.Println(err)
	}
	if result, ok := re["result"]; !ok || result != true {
		cobra.CheckErr("Submit worksheet failed!")
	}

	data := re["data"].(map[string]interface{})
	sn := data["sn"]

	return sn.(string)
}
