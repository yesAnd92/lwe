package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of lwe",
	Long:  `All software has versions. This is lwe's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v.0.0.3.beta")
	},
}
