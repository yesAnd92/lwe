package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/fileserver"
	"strconv"
)

/**
* file server
 */
var (
	fileServerPort int
	fileServerCmd  = &cobra.Command{
		Use:     `fileserver`,
		Short:   `Start a static file server`,
		Long:    `Start a static file server,designed to provide static resources, which include but are not limited to HTML files,CSS,JS,images,and videos.`,
		Example: `lwe fileserver your-file-dir [-p=8080] `,
		Args:    cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			rootDir := args[0]
			fileserver.ServerStart(strconv.Itoa(fileServerPort), rootDir)
		},
	}
)

func init() {

	fileServerCmd.PersistentFlags().IntVarP(&fileServerPort, "port", "p", 9527, "file server's web port,default 9527")
}
