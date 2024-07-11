package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/env"
	"github.com/yesAnd92/lwe/utils"
)

var envCmd = &cobra.Command{
	Use:     `env`,
	Short:   `Print all environment variable `,
	Long:    `Print all environment variable`,
	Example: `lwe env`,
	Args:    cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {

		var envV env.IEnvVariable
		op := utils.OsEnv()
		fmt.Println("op:", op)
		switch op {
		case utils.Mac:
			envV = &env.MacEVnVariable{}
		default:
			cobra.CheckErr("Not support this os!")
		}
		envInfos := envV.FindEnvInfo()
		env.EnvCat(envInfos)
	},
}
