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
		Short:   `Generate the specified file based on SQL`,
		Long:    `Generate the specified file based on SQL. Such as Java Entity,Go struct and so on`,
		Example: `lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]`,
		Args:    cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			sqlLog := args[0]
			sql, err := sqllog.ParseMybatisSqlLog(sqlLog)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("can't parse SQL: %s", sqlLog))
			}

			fmt.Println("===>")
			fmt.Println(sql)
		},
	}
)
