package gitcmd

import (
	"encoding/json"
	"fmt"
	"github.com/micro-plat/lib4go/net/http"
	"github.com/spf13/cobra"
	"lwe/utils"
	"net/url"
	"os"
	"time"
)

var (
	// CLONE_TPL git clone
	CLONE_TPL = "git clone %s %s"

	// GITLAB_GROUP_DETAIL git lab group api
	//gitlab接口文档  https://docs.gitlab.com/ee/api/groups.html#details-of-a-group
	GITLAB_GROUP_DETAIL = "%s://%s/api/v4/groups%s?private_token=%s"
)

type ProjectInfo struct {
	HttpUrlToRepo string
	Path          string
}

// CloneGroup 获取所有项目仓库
func CloneGroup(groupUrl, token, targetDir string) {

	targetDir = utils.ToAbsPath(targetDir)

	projectInfos := getRepoList(groupUrl, token)

	fmt.Printf("clone %s start...\n", groupUrl)

	for _, pj := range projectInfos {
		// 校验repo本地是否已经存在
		p := targetDir + "/" + pj.Path
		if checkPjExist(p) {
			fmt.Printf("%s  existed,skip this repository...\n", p)
			continue
		}
		cloneRepo(pj, p)
	}
	fmt.Printf("clone %s end!\n", groupUrl)
}

func cloneRepo(pj *ProjectInfo, targetDir string) {

	var cmdline = fmt.Sprintf(CLONE_TPL, pj.HttpUrlToRepo, targetDir)

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		cobra.CheckErr(result.Err())
	}

	commitMsg := result.String()
	fmt.Println(commitMsg)
}

// getRepoList 获取组下的仓库列表
func getRepoList(groupUrl, token string) []*ProjectInfo {
	u, _ := url.Parse(groupUrl)
	path := fmt.Sprintf(GITLAB_GROUP_DETAIL, u.Scheme, u.Host, u.Path, token)
	client, _ := http.NewHTTPClient()
	respBody, status, err := client.Request("get", path, "", "utf-8", nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	if status != 200 {
		fmt.Printf("error:%d\n", status)
		return nil
	}

	re := make(map[string]interface{})
	json.Unmarshal([]byte(respBody), &re)

	var projectInfos []*ProjectInfo

	projects := re["projects"].([]interface{})
	for _, project := range projects {
		projectInfos = append(projectInfos, &ProjectInfo{
			HttpUrlToRepo: project.(map[string]interface{})["http_url_to_repo"].(string),
			Path:          project.(map[string]interface{})["path"].(string),
		})
	}
	return projectInfos
}

func checkPjExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
