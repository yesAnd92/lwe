package cmd

import (
	"github.com/spf13/cobra"
	"lwe/handler/gitcmd"
)

/**
* git命令相关功能
 */
var (
	detail    bool   //是否展示每次提交变动的文件
	file      bool   //日志结果控制台输出，或者生成文件，默认控制台输出
	committer string //指定查询提交者，不指定查所有
	recentN   int8   //指定查询仓库
	start     string //开始时间
	end       string //结束时间

	gitCmd = &cobra.Command{
		Use:   "glog",
		Short: "Decrypt password of connection in .ncx file",
		Long:  `The config exported from Navicat is encrypted,ncx command can decrypt it`,
		Args:  cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			commitLogs, _ := gitcmd.GetAllGitRepoCommitLog(detail, recentN, dir, committer, start, end)

			var output gitcmd.OutputFormatter = &gitcmd.ConsoleOutput{}
			if file {
				output = &gitcmd.FileOutput{}
			}
			output.Output(commitLogs)

		},
	}
)

func init() {

	gitCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "")
	gitCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "")
	gitCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "")
	gitCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "")
	gitCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "")
	gitCmd.PersistentFlags().Int8VarP(&recentN, "recentN", "n", 10, "")
}
