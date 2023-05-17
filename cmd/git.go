package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lwe/handler/gitcmd"
)

/**
* git命令相关功能
 */
var (
	detail    bool
	committer string
	recentN   int8
	gitCmd    = &cobra.Command{
		Use:   "glog",
		Short: "Decrypt password of connection in .ncx file",
		Long:  `The config exported from Navicat is encrypted,ncx command can decrypt it`,
		Args:  cobra.MatchAll(),
		Run: func(cmd *cobra.Command, args []string) {
			commitLogs, err := gitcmd.GetCommitLog("", committer, recentN)
			if err != nil {
				fmt.Println(err)
				return
			}
			for i, log := range *commitLogs {
				fmt.Printf("#%d %s\n"+
					">>%s\n"+
					"@%s %s \n", i, log.CommitHash, log.CommitMsg, log.Username, log.CommitAt)

				if detail {
					filenames, err := gitcmd.GetChangedFile(log.CommitHash)
					if err != nil {
						fmt.Println(err)
					}
					if filenames != nil && len(filenames) > 0 {
						for _, filename := range filenames {
							fmt.Println("- " + filename)
						}
					}
				}
				fmt.Print("\n")
			}
		},
	}
)

func init() {

	gitCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "")
	gitCmd.PersistentFlags().StringVarP(&committer, "author", "a", "", "")
	gitCmd.PersistentFlags().Int8VarP(&recentN, "recentN", "n", 10, "")
}
