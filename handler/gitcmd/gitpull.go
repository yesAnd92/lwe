package gitcmd

import (
	"fmt"
	"github.com/yesAnd92/lwe/utils"
	"time"
)

func updateRepo(dir string) (string, error) {

	//指定仓库地址
	var cmdline = fmt.Sprintf(GIT_PULL, dir)

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		return "", result.Err()
	}

	reStr := result.String()
	return reStr, nil
}

// checkRepoClean 检测当前仓库是否干净
func checkRepoClean(dir string) (bool, string) {
	var cmdline = fmt.Sprintf(STATUS_CHECK_TPL, dir)

	result := utils.RunCmd(cmdline, time.Second*30)
	if result.Err() != nil {
		return false, result.Err().Error()
	}

	commitMsg := result.String()
	if len(commitMsg) == 0 {
		return true, ""
	}
	return false, commitMsg
}

// UpdateAllGitRepo 更新所有仓库
func UpdateAllGitRepo(dir string) {
	var res []string

	//相对路径转换成绝对路径进行处理
	dir = utils.ToAbsPath(dir)

	//递归找到所有的git仓库
	findGitRepo(dir, &res)

	//遍历获取每个仓库的提交信息
	for idx, gitDir := range res {
		fmt.Printf("#%d Git Repo >> %s\n", idx+1, gitDir)
		if clean, msg := checkRepoClean(gitDir); !clean {
			//存在未提交的变动，防止冲突等问题，不进行更新
			fmt.Println(msg)
			fmt.Println(">> Modified Files in this Repo have not been submitted yet, terminating the update!")
			fmt.Println()
			continue
		}

		result, err := updateRepo(gitDir)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
		fmt.Println()
	}

}
