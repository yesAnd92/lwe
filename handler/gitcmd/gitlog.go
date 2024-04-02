package gitcmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// CommitLog 提交记录
type CommitLog struct {
	Branch       string   //Branch
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
func GetCommitLog(detail bool, recentN int16, dir, author, start, end string, branchs bool) (*[]CommitLog, error) {

	var logs []CommitLog

	branchInfo := ListRepoAllBranch(dir)

	//filter given branch
	if !branchs {
		branchInfo.branchs = []string{branchInfo.curr}
	}

	// find all branch commit log
	for _, branch := range branchInfo.branchs {

		cmdline := buildCmdline(dir, branch, recentN, author, start, end)

		result := utils.RunCmd(cmdline, time.Second*30)
		if result.Err() != nil {
			return nil, result.Err()
		}

		reStr := result.String()
		if len(reStr) == 0 {
			continue
		}

		commitLines := strings.Split(reStr, "\n")
		for _, msg := range commitLines {
			//win环境下会多“'”符号，替换去除
			msg = strings.Trim(msg, "'")
			infoArr := strings.Split(msg, "*-*")
			//commitAt
			commitAtMill, _ := strconv.ParseInt(infoArr[2], 10, 64)

			log := CommitLog{
				Branch:     branch,
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
	}

	//merge same commit log in different branch
	mergeLog := mergeAndSortCommitLog(&logs, int(recentN))

	return &mergeLog, nil
}

func mergeAndSortCommitLog(logs *[]CommitLog, n int) []CommitLog {

	var uniqueLogs []CommitLog
	seen := make(map[string]struct{})
	for _, log := range *logs {
		if _, ok := seen[log.CommitHash]; !ok {
			seen[log.CommitHash] = struct{}{}
			uniqueLogs = append(uniqueLogs, log)
		}
	}

	sort.Slice(uniqueLogs, func(i, j int) bool {
		return uniqueLogs[i].CommitAt > uniqueLogs[j].CommitAt
	})

	if len(uniqueLogs) < n {
		n = len(uniqueLogs)
	}

	return uniqueLogs[:n]
}

func buildCmdline(dir string, branch string, recentN int16, author string, start string, end string) string {
	//使用bytes.Buffer这种方式拼接字符串会%!h(MISSING)？
	//指定仓库地址
	var cmdline = fmt.Sprintf(LOG_TPL, dir, branch)

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
	return cmdline
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
func GetAllGitRepoCommitLog(detail bool, recentN int16, dir, author, start, end string, branchs bool) (*[]ResultLog, error) {
	var res []string
	var reLog []ResultLog

	//相对路径转换成绝对路径进行处理
	dir = utils.ToAbsPath(dir)

	//递归找到所有的git仓库
	findGitRepo(dir, &res)

	//遍历获取每个仓库的提交信息
	for _, gitDir := range res {
		commitLogs, err := GetCommitLog(detail, recentN, gitDir, author, start, end, branchs)
		if err != nil {
			cobra.CheckErr(err)
		}
		if len(*commitLogs) > 0 {
			reLog = append(reLog, ResultLog{
				RepoName:   gitDir,
				CommitLogs: commitLogs,
			})
		}
	}

	return &reLog, nil
}
