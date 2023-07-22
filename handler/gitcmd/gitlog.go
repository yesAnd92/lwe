package gitcmd

import (
	"errors"
	"fmt"
	"github.com/yesAnd92/lwe/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (

	//git log
	LOG_TPL            = "git -C %s --no-pager log  --no-merges "
	LOG_FORMAT_TPL     = `--format=format:'%h*-*%an*-*%ct*-*%s' ` //使用*-*作为分隔符
	LOG_AUTHOR_TPL     = `--author=%s `
	LOG_START_DATE_TPL = `--since=%s `
	LOG_END_DATE_TPL   = `--until=%s `
	LOG_RECENTN_TPL    = `-n %d `

	//git show
	SHOW_TPL = "git show %s"
)

// CommitLog 提交记录
type CommitLog struct {
	CommitHash   string   //abbreviated commit hash
	Username     string   //username
	CommitAt     string   //commit time
	CommitMsg    string   //commit log
	FilesChanged []string //changed file arr
}

// ResultLog 封装日志分析结果
type ResultLog struct {
	RepoName   string //git repository name
	CommitLogs *[]CommitLog
}

// GetCommitLog 获取提交日志
func GetCommitLog(detail bool, recentN int16, dir, author, start, end string) (*[]CommitLog, error) {

	//使用bytes.Buffer这种方式拼接字符串会%!h(MISSING)？
	//指定仓库地址
	var cmdline = fmt.Sprintf(LOG_TPL, dir)

	if recentN >= 0 {
		cmdline += fmt.Sprintf(LOG_RECENTN_TPL, recentN)
	}

	cmdline += LOG_FORMAT_TPL

	if len(author) > 0 {
		cmdline += fmt.Sprintf(LOG_AUTHOR_TPL, author)
	}

	if len(start) > 0 {
		cmdline += fmt.Sprintf(LOG_START_DATE_TPL, start)
	}

	if len(end) > 0 {
		cmdline += fmt.Sprintf(LOG_END_DATE_TPL, end)
	}

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var logs []CommitLog
	reStr := result.String()
	if len(reStr) == 0 {
		return nil, errors.New("no matching commit log found！")
	}

	commitLines := strings.Split(reStr, "\n")
	for _, msg := range commitLines {
		//win环境下会多“'”符号，替换去除
		msg = strings.Trim(msg, "'")
		infoArr := strings.Split(msg, "*-*")
		//commitAt
		commitAtMill, _ := strconv.ParseInt(infoArr[2], 10, 64)

		log := CommitLog{
			CommitHash: infoArr[0],
			Username:   infoArr[1],
			CommitAt:   time.UnixMilli(commitAtMill * 1000).Format("2006-01-02 15:04:05"),
			CommitMsg:  infoArr[3],
		}

		//change file
		if detail {
			filesChanged, err := GetChangedFile(infoArr[0])
			if err == nil {
				log.FilesChanged = filesChanged
			}
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

// GetAllGitRepoCommitLog 封装所有仓库的提交信息
func GetAllGitRepoCommitLog(detail bool, recentN int16, dir, author, start, end string) (*[]ResultLog, error) {
	var res []string
	var reLog []ResultLog

	//相对路径转换成绝对路径进行处理
	dir = utils.ToAbsPath(dir)

	//递归找到所有的git仓库
	findGitRepo(dir, &res)

	//遍历获取每个仓库的提交信息
	for _, gitDir := range res {
		commitLogs, err := GetCommitLog(detail, recentN, gitDir, author, start, end)
		if err == nil {
			reLog = append(reLog, ResultLog{
				RepoName:   gitDir,
				CommitLogs: commitLogs,
			})
		}
	}

	return &reLog, nil
}
