package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/gitcmd"
)

/**
* git命令相关功能
 */
var (
	detail    = false //是否展示每次提交变动的文件 !!!没找到合适的展示方式，且性能稳定性不可控，暂时先不开放
	file      bool    //日志结果控制台输出，或者生成文件，默认控制台输出
	committer string  //指定查询提交者，不指定查所有
	recentN   int16   //指定查询仓库数量
	start     string  //开始时间
	end       string  //结束时间

	//gcl
	token string //克隆所需要的token

	glogCmd = &cobra.Command{
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

	glCmd = &cobra.Command{
		Use:   "gl",
		Short: "Update all git repository under the given dir ",
		Long:  `Update all git repository under the given dir ,the repository that has modified files will not be updated!`,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			gitcmd.UpdateAllGitRepo(dir)

		},
	}

	gclCmd = &cobra.Command{
		Use:   "gcl",
		Short: "Git clone all git repository under the given git group ",
		Long:  `Git clone all git repository under the given git group `,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			if len(token) == 0 {
				cobra.CheckErr("please confirm token is not empty!")
			}
			var dir = "."
			if len(args) > 1 {
				dir = args[1]
			}

			gitcmd.CloneGroup(args[0], token, dir)

		},
	}

	gstCmd = &cobra.Command{
		Use:   "gst",
		Short: "Get all git repository status under the given dir ",
		Long:  `Get all git repository status under the given dir `,
		Args:  cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			gitcmd.GetAllGitRepoStatus(dir)

		},
	}
)

func init() {

	//gitCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "")
	glogCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "result output to file,default value is false (meaning output to console). ")
	glogCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "specify name of committer ")
	glogCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "specify the start of commit date. eg.'yyyy-MM-dd'")
	glogCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "specify the end of commit date. eg.'yyyy-MM-dd'")
	glogCmd.PersistentFlags().Int16VarP(&recentN, "recentN", "n", 10, "specify the number of commit log for each git repo.")

	//gcl
	gclCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "private token")
}
