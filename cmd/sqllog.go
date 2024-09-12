package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/sqllog"
)

/**
* mybatis sql log parse
 */

var (
	sqlLogCmd = &cobra.Command{
		Use:     `sqllog`,
		Short:   `Parse mybatis sql logs and fill placeholders with parameters`,
		Long:    `Copy mybatis sql log ,extract sql info and fill placeholders with parameters`,
		Example: `lwe sqllog 'mybatis sql log'`,
		Args:    cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			sqlLog := args[0]

			sql, err := sqllog.ParseMybatisSqlLog(sqlLog)
			if err != nil {
				cobra.CheckErr(err.Error())
			}

			fmt.Println("======>")

			fmt.Println(sql)
		},
	}
)
