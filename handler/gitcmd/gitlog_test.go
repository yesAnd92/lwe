package gitcmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"
)

func TestGetCommit(t *testing.T) {

	commitLogs, err := GetCommitLog("yesand", 20)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, log := range *commitLogs {
		fmt.Printf("-----------Hash:%s-----------\n", log.CommitHash)
		fmt.Printf("@%s  %s\n"+
			"commit msg: %s\n\n", log.Username, log.CommitAt, log.CommitMsg)
	}
}

func TestGetChangedFile(t *testing.T) {
	filenames, err := GetChangedFile("6f635d7")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, filename := range filenames {
		fmt.Println(filename)
	}
}

func TestGetChangedDir(t *testing.T) {

	findPath := "/Users/wangyj/ideaProject"
	var res []string
	findGitRepo(findPath, &res)
	for _, s := range res {
		fmt.Println(s)

	}
}

func findGitRepo(dir string, res *[]string) {
	var files []string
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		//return files, err
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
	for _, fName := range files {
		findGitRepo(path.Join(dir, fName), res)
	}
	//return files, nil
	//fmt.Println(files)
}
