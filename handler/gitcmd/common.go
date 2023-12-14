package gitcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/utils"
	"os"
	"path/filepath"
	"time"
)

func findGitRepo(dir string, res *[]string) {

	// Search for git repo dir recursively
	nextGitRepo(dir, res)

	//not fund git repo under dir
	//check the dir is under git repo
	if result := utils.RunCmd(EXIST_GIT_REPO, time.Second*10); len(*res) == 0 && result.String() == "true" {

		//if result.Err() != nil {
		//	cobra.CheckErr(result.Err())
		//}

		fmt.Println(">>>>", result.String())
		*res = append(*res, dir)
	}
}

func nextGitRepo(dir string, res *[]string) {
	var files []string
	fileInfo, err := os.ReadDir(dir)
	if err != nil {
		cobra.CheckErr(fmt.Errorf(" The dir '%s' is not exist!\n", dir))
		return
	}

	for _, file := range fileInfo {
		//当前目录是git仓库，没必要继续遍历
		if ".git" == file.Name() {
			*res = append(*res, dir)
			return
		}
		if file.IsDir() {
			files = append(files, file.Name())
		}
	}

	//目录下的子目录递归遍历
	for _, fName := range files {
		nextGitRepo(filepath.Join(dir, fName), res)
	}
}
