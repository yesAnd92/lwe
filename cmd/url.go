package cmd

import "github.com/spf13/cobra"
import "github.com/yesAnd92/lwe/handler/url"

var urlCmd = &cobra.Command{
	Use:     `url`,
	Short:   `Format request url to increase readability`,
	Long:    `Format request url to increase readability`,
	Example: `lwe url yourUrl`,
	Args:    cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		url.HandleUrlPathParams(args[0])
	},
}
