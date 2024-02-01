package gitcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
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

type branchInfo struct {
	curr    string
	branchs []string
}

// ListRepoAllBranch list all git Branch under repository
func ListRepoAllBranch(repo string) (re *branchInfo) {

	// get current dir
	originalDir, _ := os.Getwd()

	err := os.Chdir(repo)
	if err != nil {
		log.Fatal(err)
	}

	// change to original dir
	defer os.Chdir(originalDir)

	var curr string
	var branchs []string

	if result := utils.RunCmd(GIT_BRANCH, time.Second*10); result.Err() == nil {

		branchSplit := strings.Split(result.String(), "\n")
		for _, branch := range branchSplit {
			branch = strings.TrimSpace(branch)
			//  current Branch mark with '*'
			if strings.HasPrefix(branch, "*") {
				branch = strings.ReplaceAll(branch, "* ", "")
				curr = branch
			}
			branchs = append(branchs, branch)

		}

		re = &branchInfo{
			curr:    curr,
			branchs: branchs,
		}
	}
	return
}
