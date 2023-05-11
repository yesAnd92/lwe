package gitcmd

import (
	"errors"
	"fmt"
	"lwe/utils"
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
	LOG_FORMAT_TPL  = `--format=format:'%h^*%an^*%ct^*%s' ` //使用^*作为分隔符
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
func GetCommitLog(author string, recentN int8) (*[]CommitLog, error) {
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
		infoArr := strings.Split(msg, "^*")
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
