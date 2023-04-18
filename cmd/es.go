package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xwb1989/sqlparser"
	"lwe/handler/es"
)

/**
* es命令
* 将sql语句转译成dsl语句
 */
var (
	fmtPretty bool
	esCmd     = &cobra.Command{
		Use:   "es",
		Short: "Translate SQL to elasticsearch's DSL",
		Long:  `Translate SQL to elasticsearch's DSL`,
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			sql := args[0]
			//使用sqlparse对SQL进行解析
			stmt, err := sqlparser.Parse(sql)
			if err != nil {
				fmt.Println("Something error in your sql:", err)
				fmt.Println("Please re-check syntax and try it again!")
				return
			}

			var dsl, esType, urlMethod string
			switch stmt.(type) {
			case *sqlparser.Select:
				dsl, esType, err = es.HandleSelect(stmt.(*sqlparser.Select))
				urlMethod = fmt.Sprintf("POST /%s/_search", esType)
			case *sqlparser.Delete:
				fmt.Println("Delete syntax is not supported this version!")
				return
			case *sqlparser.Update:
				fmt.Println("Update syntax is not supported this version!")
				return
			default:
				fmt.Println("This syntax is supported this version!")
				return
			}

			if err != nil {
				fmt.Println(err)
				return
			}
			//输出请求路径，比如 POST /index_type/_search
			fmt.Println(urlMethod)

			if fmtPretty {
				//需要美化
				var re map[string]interface{}
				json.Unmarshal([]byte(dsl), &re)
				pr, _ := json.MarshalIndent(re, "", "  ")
				dsl = string(pr)
			}

			fmt.Println(dsl)
		},
	}
)

func init() {

	esCmd.PersistentFlags().BoolVarP(&fmtPretty, "pretty", "p", false, "Beautify DSL")
}
