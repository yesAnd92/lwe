package gitcmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

const REPORT_PATH = "./commit.log"

//输出形式：控制台数据、写入文件

type OutputFormatter interface {
	Output(*[]ResultLog)
}

type ConsoleOutput struct {
}

func (c *ConsoleOutput) Output(resLogs *[]ResultLog) {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	// set column width
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Number:   4,
			WidthMax: 100,
			WidthMaxEnforcer: func(col string, maxLen int) string {

				return text.WrapText(col, maxLen)
			},
		},
	})
	t.AppendHeader(table.Row{"Branch", "Hash", "Author", "Commit", "Time"})

	if *resLogs == nil {
		fmt.Printf("No matching commit log found in this Dir\n")
		return
	}

	for idx, res := range *resLogs {
		logs := res.CommitLogs
		fmt.Printf("#%d Git Repo >> %s\n", idx+1, res.RepoName)

		for _, log := range *logs {
			t.AppendRow(table.Row{log.Branch, log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
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

	if *resLogs == nil {
		commitData.WriteString("No matching commit log found in this git repo\n")
	}

	for idx, res := range *resLogs {
		logs := res.CommitLogs
		commitData.WriteString(fmt.Sprintf("#%d Git Repo >> %s\n", idx+1, res.RepoName))

		for _, log := range *logs {
			t.AppendRow(table.Row{log.CommitHash, log.Username, log.CommitMsg, log.CommitAt})
			if len(log.FilesChanged) > 0 {
				t.AppendRow(table.Row{log.FilesChanged})
			}
		}

		t.Render()
		t.ResetRows()
		commitData.WriteString("\n")
	}

	path, _ := filepath.Abs(REPORT_PATH)
	f, err := os.Create(path)
	if f != nil {
		defer f.Close()
	}

	if err != nil {
		cobra.CheckErr(err)
	}
	f.Write(commitData.Bytes())

	fmt.Println("Commit log has finished >> " + path)
}
