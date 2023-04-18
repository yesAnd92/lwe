package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	md5Cmd = &cobra.Command{
		Use:   "md5",
		Short: "Get a md5 for the given value or  a random md5 value",
		Long:  `Get a md5 for the given value. If not specify value ,it will give a random md5 value`,
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			specifyValue := args[0]
			sum := md5.Sum([]byte(specifyValue))
			fmt.Println(hex.EncodeToString(sum[:]))
		},
	}
)
