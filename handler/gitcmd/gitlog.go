package gitcmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"lwe/utils"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	//userinfo
	USERNAME_TPL = ""

	//git log
	LOG_TPL         = "git --no-pager log  --no-merges "
	LOG_FORMAT_TPL  = `--format=format:'%h**%an**%ct**%s' ` //使用^*作为分隔符
	LOG_AUTHOR_TPL  = `--author=%s `
	LOG_RECENTN_TPL = `-n %d `

	//git show
	SHOW_TPL = "git show %s"
)

type CommitLog struct {
	CommitHash string //abbreviated commit hash
	Username   string //username
	CommitAt   string //commit time
	CommitMsg  string //commit log
}

// GetCommitLog 获取提交日志
func GetCommitLog(dir string, author string, recentN int8) (*[]CommitLog, error) {

	if len(dir) == 0 {
		//指定了目录，切换到指定目录执行命令
		if err := os.Chdir(dir); err != nil {
			log.Fatal(err)
		}
	}

	//使用bytes.Buffer这种方式拼接字符串会%!h(MISSING)？
	var cmdline = LOG_TPL
	if recentN >= 0 {
		cmdline += fmt.Sprintf(LOG_RECENTN_TPL, recentN)
	}

	cmdline += LOG_FORMAT_TPL

	if len(author) > 0 {
		cmdline += fmt.Sprintf(LOG_AUTHOR_TPL, author)
	}

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var logs []CommitLog
	if len(result.String()) == 0 {
		return nil, errors.New("no matching commit log found！")
	}
	commitLines := strings.Split(result.String(), "\n")
	for _, msg := range commitLines {
		infoArr := strings.Split(msg, "**")
		//commitAt
		commitAtMill, _ := strconv.ParseInt(infoArr[2], 10, 64)
		log := CommitLog{
			CommitHash: infoArr[0],
			Username:   infoArr[1],
			CommitAt:   time.UnixMilli(commitAtMill * 1000).Format("2006-01-02 15:04:05"),
			CommitMsg:  infoArr[3],
		}
		logs = append(logs, log)
	}
	return &logs, nil
}

// GetChangedFile 获取本次提交变动的文件名
func GetChangedFile(commitId string) ([]string, error) {
	var cmdline = fmt.Sprintf(SHOW_TPL, commitId)

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		return nil, result.Err()
	}

	commitMsg := result.String()
	if len(commitMsg) == 0 {
		return nil, errors.New("")
	}
	var fileNames []string
	re := regexp.MustCompile("--- a.+")
	finds := re.FindAllString(commitMsg, -1)
	for _, find := range finds {
		find = find[strings.LastIndex(find, "/")+1:]
		fileNames = append(fileNames, find)
	}
	return fileNames, nil
}

func GetAllGitRepoCommitLog(dir string) map[string]*[]CommitLog {
	var res []string
	repoCommitMap := make(map[string]*[]CommitLog)
	findGitRepo(dir, &res)
	for _, gitDir := range res {
		commitLog, err := GetCommitLog(gitDir, "", 10)
		if err == nil {
			repoCommitMap[gitDir] = commitLog
		}
	}
	return repoCommitMap

}

func findGitRepo(dir string, res *[]string) {
	var files []string
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic("dir is wrong: " + dir)
	}

	for _, file := range fileInfo {
		//当前目录是git仓库，没必要继续遍历
		if ".git" == file.Name() {
			//fmt.Println(dir)
			*res = append(*res, dir)
			return
		}
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}

	//目录下的子目录递归遍历
	for _, fName := range files {
		findGitRepo(path.Join(dir, fName), res)
	}
}
