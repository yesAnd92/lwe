package gitcmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
)

func findGitRepo(dir string, res *[]string) {
	var files []string
	fileInfo, err := ioutil.ReadDir(dir)
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
		findGitRepo(filepath.Join(dir, fName), res)
	}
}
