package cmd

import "github.com/spf13/cobra"
import "github.com/yesAnd92/lwe/handler/url"

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "format request url to increase readability",
	Long:  `format request url to increase readability`,
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		url.HandleUrlPathParams(args[0])
	},
}
