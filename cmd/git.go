package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/gitcmd"
	"github.com/yesAnd92/lwe/utils"
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
	branchs   bool    //show all branch,default show current branch

	//gcl
	token string //克隆所需要的token

	glogCmd = &cobra.Command{
		Use:     `glog`,
		Short:   `Get all git repository commit log under the given dir `,
		Long:    `Get all git repository commit log under the given dir ,and  specify author，date etc. supported!`,
		Example: `lwe glog [git repo dir] [-a=yesAnd] [-n=50] [-s=2023-08-04] [-e=2023-08-04] [-b=ture]`,
		Args:    cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			if recentN > int16(500) {
				cobra.CheckErr("recentN can't exceed 500")
			}

			commitLogs, err := gitcmd.GetAllGitRepoCommitLog(detail, recentN, dir, committer, start, end, branchs)
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
		Use:     `gl`,
		Short:   `Update all git repository under the given dir `,
		Long:    `Update all git repository under the given dir ,the repository that has modified files will not be updated!`,
		Example: `lwe gl [git repo dir]`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			gitcmd.UpdateAllGitRepo(dir)

		},
	}

	gclCmd = &cobra.Command{
		Use:     `gcl`,
		Short:   `Git clone all git repository under the given git group `,
		Long:    `Git clone all git repository under the given git group `,
		Example: `lwe gcl gitGroupUrl [dir for this git group] -t=yourToken`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(1)),
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
		Use:     `gst`,
		Short:   `Get all git repository status under the given dir `,
		Long:    `Get all git repository status under the given dir `,
		Example: `lwe gst [your git repo dir]`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			gitcmd.GetAllGitRepoStatus(dir)

		},
	}

	gSumCmd = &cobra.Command{
		Use:     `gsum`,
		Short:   `Summary git log with AI's help' `,
		Long:    `With the help of AI, merge similar or streamline Git commit logs.`,
		Example: `lwe gsum [git repo dir] [-a=yesAnd] [-s=2023-08-04] [-e=2023-08-04]`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(args) > 0 {
				dir = args[0]
			}

			gitcmd.GitLogSummary(detail, dir, committer, start, end)

		},
	}

	gcmsgCmd = &cobra.Command{
		Use:     `gcmsg`,
		Short:   `Generate commit msg with AI's help `,
		Long:    `Generate commit msg with AI's help`,
		Example: `lwe gcmsg`,
		Args:    cobra.MatchAll(cobra.MinimumNArgs(0)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			//trans to abslute path
			dir = utils.ToAbsPath(dir)

			//git commit msg from ai
			commit := gitcmd.GetGitCommitMsg(dir)

			//push to origin repo
			gitcmd.CommitAndPush(dir, commit)

		},
	}
)

func init() {

	glogCmd.PersistentFlags().BoolVarP(&file, "file", "f", false, "result output to file,default value is false (meaning output to console). ")
	glogCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "specify name of committer ")
	glogCmd.PersistentFlags().StringVarP(&start, "start", "s", "", "specify the start of commit date. eg.'yyyy-MM-dd'")
	glogCmd.PersistentFlags().StringVarP(&end, "end", "e", "", "specify the end of commit date. eg.'yyyy-MM-dd'")
	glogCmd.PersistentFlags().Int16VarP(&recentN, "recentN", "n", 10, "specify the number of commit log for each git repo. Limit 500 ")
	glogCmd.PersistentFlags().BoolVarP(&branchs, "branchs", "b", false, "show all branch logs,default is false (meaning show current branch). ")

	//gcl
	gclCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "private token")
}
