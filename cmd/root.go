package cmd

import (
	"bytes"
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

// ExecuteCommand executes the  command
func executeCommand(cmd *cobra.Command, args ...string) (outBf *bytes.Buffer, err error) {
	root := rootCmd
	root.AddCommand(cmd)

	outBf = new(bytes.Buffer)
	root.SetOut(outBf)
	root.SetErr(outBf)
	root.SetArgs(args)

	err = root.Execute()
	if err != nil {
		return nil, err
	}
	return outBf, err
}

func init() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(fmtCmd)
	rootCmd.AddCommand(md5Cmd)
	rootCmd.AddCommand(esCmd)
	rootCmd.AddCommand(navicatCmd)
	rootCmd.AddCommand(urlCmd)
	rootCmd.AddCommand(glogCmd)
	rootCmd.AddCommand(glCmd)
	rootCmd.AddCommand(gclCmd)
	rootCmd.AddCommand(gstCmd)
	rootCmd.AddCommand(fsyncCmd)
	rootCmd.AddCommand(fileServerCmd)
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(sqlLogCmd)
	rootCmd.AddCommand(gSumCmd)
	rootCmd.AddCommand(gcmsgCmd)

}
