package gitcmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"testing"
)

func TestGetCommit(t *testing.T) {

	commitLogs, err := GetCommitLog(false, 10, ".", "yesand", "", "")
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

func TestGetAllGitRepoCommitLog(t *testing.T) {

	//resLogs, _ := GetAllGitRepoCommitLog(false, 10, "/Users/wangyj/ideaProject/work/", "wangyj", "2023-05-01", "")
	resLogs, _ := GetAllGitRepoCommitLog(false, 10, "D:\\ideaProject\\my", "", "", "")

	//控制台
	console := ConsoleOutput{}
	console.Output(resLogs)

	//写文件
	file := FileOutput{}
	file.Output(resLogs)

}

func TestTable(t *testing.T) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Author", "Commit", "Time"})

	for i := 0; i < 10; i++ {
		fmt.Printf("Git Repo >> %s\n", "aaaa")
		for j := 0; j < 10; j++ {
			itoa := strconv.Itoa(j)
			table.Append([]string{"我那个" + itoa, "a" + itoa, "撒旦法师的" + itoa, "阿斯蒂芬" + itoa + "\n"})
		}

		table.Render()
		table.ClearRows()
		fmt.Println()
	}

}
