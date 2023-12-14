package gitcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/utils"
	"time"
)

// checkRepoClean 查看当前仓库状态
func printRepoStatus(dir string) {
	var cmdline = fmt.Sprintf(STATUS_TPL, dir)

	result := utils.RunCmd(cmdline, time.Second*5)
	if result.Err() != nil {
		cobra.CheckErr(result.Err().Error())
	}

	fmt.Println(result.String())
}

// GetAllGitRepoStatus 查看所有仓库的状态
func GetAllGitRepoStatus(dir string) {
	var res []string

	//相对路径转换成绝对路径进行处理
	dir = utils.ToAbsPath(dir)

	//递归找到所有的git仓库
	findGitRepo(dir, &res)

	//遍历获取每个仓库的提交信息
	for idx, gitDir := range res {
		fmt.Printf("#%d Git Repo >> %s\n", idx+1, gitDir)
		printRepoStatus(gitDir)
		fmt.Println()
	}

}
