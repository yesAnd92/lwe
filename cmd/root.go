package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lwe",
		Short: "meaning leave work early!",
		Long:  `meaning leave work early!`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(sqlCmd)
	rootCmd.AddCommand(md5Cmd)
	rootCmd.AddCommand(esCmd)
	rootCmd.AddCommand(navicatCmd)
	rootCmd.AddCommand(glogCmd)
	rootCmd.AddCommand(gplCmd)
	rootCmd.AddCommand(urlCmd)
}
