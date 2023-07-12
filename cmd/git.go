package cmd

import (
	"github.com/spf13/cobra"
	"lwe/handler/gitcmd"
)

/**
* git命令相关功能
 */
var (
	detail    = false //是否展示每次提交变动的文件 !!!没找到合适的展示方式，且性能稳定性不可控，暂时先不开放
	file      bool    //日志结果控制台输出，或者生成文件，默认控制台输出
	committer string  //指定查询提交者，不指定查所有
	recentN   int16   //指定查询仓库
	start     string  //开始时间
	end       string  //结束时间

	gitCmd = &cobra.Command{
		Use:   "glog",
		Short: "Get all git repository commit log under the given dir ",
		Long:  `Get all git repository commit log under the given dir ,and  specify author，date etc. supported!`,
		Args:  cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			if recentN > int16(200) {
				cobra.CheckErr("recentN can't exceed 200")
			}

			commitLogs, err := gitcmd.GetAllGitRepoCommitLog(detail, recentN, dir, committer, start, end)
			if err != nil {
				cobra.CheckErr(err)
			}

			var output gitcmd.OutputFormatter = &gitcmd.ConsoleOutput{}
			if file {
				output = &gitcmd.FileOutput{}
			}
			output.Output(commitLogs)

		},
	}
)

func init() {

	gitCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "result output to file,default value is false (meaning output to console). ")
	gitCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "specify name of committer ")
	gitCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "specify the start of commit date. eg.'yyyy-MM-dd'")
	gitCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "specify the end of commit date. eg.'yyyy-MM-dd'")
	gitCmd.PersistentFlags().Int16VarP(&recentN, "recentN", "n", 10, "specify the number of commit log for each git repo.")
}
