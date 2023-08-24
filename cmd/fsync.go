package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/sync"
)

var (
	dryrRun bool

	fsyncCmd = &cobra.Command{
		Use:     `fsync`,
		Short:   `Sync file from source dir to target dir`,
		Long:    `Sync file from source dir to target dir,and it will skip same name file`,
		Example: `lwe fsync sourceDir targetDir [-d=true]`,
		Args:    cobra.MatchAll(cobra.ExactArgs(2)),
		Run: func(cmd *cobra.Command, args []string) {
			sourceDir := args[0]
			targetDir := args[1]

			var thenDo sync.CompareThenDoIfa = &sync.CopyCompareThenDo{}
			if dryrRun {
				thenDo = &sync.DisplayCompareThenDo{}
			}
			fsync := sync.InitFsync(sourceDir, targetDir)

			//compare source and target dir diff
			fsync.DiffDir()

			fsync.Sync(thenDo)
		},
	}
)

func init() {

	//dry-run
	fsyncCmd.PersistentFlags().BoolVarP(&dryrRun, "dry-run", "d", false, "Because fsync can make some significant changes, you might prefer to add --dry-run=true option"+
		" to the command line to preview what fsync plans to do")
}
