package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/dp"
)

/**
*dp命令相关功能
 */
var (

	//gdp命令
	path      string //指定目录
	worksheet bool   //是否需要提交工单，默认true
	//gdp部署命令
	dpCmd = &cobra.Command{
		Use:   "dp",
		Short: "升级部署助手",
		Long:  `升级部署助手，可以根据参数生成tag并推送到git仓库，自动填写工单，生成升级脚本`,
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		Run: func(cmd *cobra.Command, args []string) {

			var dir = "."
			if len(path) > 0 {
				dir = path
			}

			var updateModule = args[0]

			var msg = args[1]

			//构造部署信息
			dpObj := dp.BuildDdpInfo(updateModule, dir, msg)

			//推送tag到远程仓库
			dp.CommitAndPushTag(dpObj.Tag, path)

			//需要工单提交
			if worksheet {
				//登录 itsm
				cookieJar := dp.LoginItsm(dpObj)

				//提交工单
				dp.SubmitWorkSheet(cookieJar, dpObj)
			}

			//输出升级部署命令
			dp.PrintTplResult(dpObj)

		},
	}
)

func init() {

	dpCmd.Flags().StringVarP(&path, "path", "p", "", "specify deploy ")
	dpCmd.Flags().BoolVarP(&worksheet, "worksheet", "w", true, " whether submit worksheet")
}
