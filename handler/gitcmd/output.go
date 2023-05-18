package gitcmd

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"path/filepath"
)

const REPORT_PATH = "./commit-log-report.txt"

//输出形式：控制台数据、写入文件

type OutputFormatter interface {
	Output(*[]ResultLog)
}

type ConsoleOutput struct {
}

func (c *ConsoleOutput) Output(resLogs *[]ResultLog) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hash", "Author", "Commit", "Time"})

	for _, res := range *resLogs {
		logs := res.CommitLogs
		fmt.Printf("Git Repo >> %s\n", res.RepoName)
		for _, log := range *logs {
			table.Append([]string{log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
		}

		table.Render()
		table.ClearRows()
		fmt.Println()
	}

}

type FileOutput struct {
}

func (c *FileOutput) Output(resLogs *[]ResultLog) {

	//渲染的分析日志放到buffer中，最后一起写入文件
	commitData := &bytes.Buffer{}
	table := tablewriter.NewWriter(commitData)
	table.SetHeader([]string{"Hash", "Author", "Commit", "Time"})

	for _, res := range *resLogs {
		logs := res.CommitLogs
		commitData.WriteString(fmt.Sprintf("Git Repo >> %s\n", res.RepoName))
		for _, log := range *logs {
			table.Append([]string{log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
		}

		table.Render()
		table.ClearRows()
		commitData.WriteString("\n")
	}

	path, _ := filepath.Abs(REPORT_PATH)
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Println("Create go file err", err)
		return
	}
	f.Write(commitData.Bytes())

	fmt.Println("Finished Report,path: " + path)
}
