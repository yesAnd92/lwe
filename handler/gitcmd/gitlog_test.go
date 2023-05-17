package gitcmd

import (
	"fmt"
	"testing"
)

func TestGetCommit(t *testing.T) {

	commitLogs, err := GetCommitLog("", "yesand", 20)
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

//func TestGetChangedDir(t *testing.T) {
//	beforePwd, _ := os.Getwd()
//	fmt.Println("切换前工作目录:", beforePwd)
//	execDir := "D:\\ideaProject\\my\\go_works"
//
//	if err := os.Chdir(execDir); err != nil {
//		log.Fatal(err)
//	}
//	pwd, _ := os.Getwd()
//	fmt.Println("切换后工作目录:", pwd)
//
//}

func TestGetAllGitRepoCommitLog(t *testing.T) {

	resMap := GetAllGitRepoCommitLog("D:\\ideaProject\\my\\go_works")
	for k, commitLogs := range resMap {
		fmt.Println(k)
		for i, log := range *commitLogs {
			fmt.Printf("#%d %s\n"+
				">>%s\n"+
				"@%s %s \n", i, log.CommitHash, log.CommitMsg, log.Username, log.CommitAt)

			//if detail {
			//	filenames, err := gitcmd.GetChangedFile(log.CommitHash)
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//	if filenames != nil && len(filenames) > 0 {
			//		for _, filename := range filenames {
			//			fmt.Println("- " + filename)
			//		}
			//	}
			//}
			fmt.Print("\n")
		}

	}

}
