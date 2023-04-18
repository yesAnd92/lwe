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
		fmt.Println("0.1beta")
	},
}

func init() {

	//versionCmd.PersistentFlags().StringP("abc", "c", "YOUR NAME", "author name for copyright attribution")
	//versionCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	//rootCmd.AddCommand(versionCmd)
}
