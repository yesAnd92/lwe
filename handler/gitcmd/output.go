package gitcmd

import (
	"bytes"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
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

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Hash", "Author", "Commit", "Time"})

	for idx, res := range *resLogs {
		logs := res.CommitLogs
		fmt.Printf("#%d Git Repo >> %s\n", idx+1, res.RepoName)
		for _, log := range *logs {
			t.AppendRow(table.Row{log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
		}

		t.Render()
		t.ResetRows()
		fmt.Println()
	}
}

type FileOutput struct {
}

func (c *FileOutput) Output(resLogs *[]ResultLog) {

	//渲染的分析日志放到buffer中，最后一起写入文件
	commitData := &bytes.Buffer{}
	t := table.NewWriter()
	t.SetOutputMirror(commitData)
	t.AppendHeader(table.Row{"Hash", "Author", "Commit", "Time"})

	for idx, res := range *resLogs {
		logs := res.CommitLogs
		commitData.WriteString(fmt.Sprintf("#%d Git Repo >> %s\n", idx+1, res.RepoName))
		for _, log := range *logs {
			t.AppendRow(table.Row{log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
		}

		t.Render()
		t.ResetRows()
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
