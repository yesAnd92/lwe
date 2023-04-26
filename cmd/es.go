package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xwb1989/sqlparser"
	"lwe/handler/es"
)

const (
	CURL_TPL = `curl -XPOST -H "Content-Type: application/json" -u {username}:{password} {ip:port}/%s/_search?pretty -d '%s' `
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

			var dsl, esType string
			switch stmt.(type) {
			case *sqlparser.Select:
				dsl, esType, err = es.HandleSelect(stmt.(*sqlparser.Select))
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

			if fmtPretty {
				//需要美化
				var re map[string]interface{}
				json.Unmarshal([]byte(dsl), &re)
				pr, _ := json.MarshalIndent(re, "", "  ")
				dsl = string(pr)
			}
			fmt.Printf(CURL_TPL, esType, dsl)
		},
	}
)

func init() {

	esCmd.PersistentFlags().BoolVarP(&fmtPretty, "pretty", "p", false, "Beautify DSL")
	//esCmd.PersistentFlags().BoolVarP(&fmtPretty, "password", "pwd", false, "Generate curl with user and password")
}

//curl -u elastic:vvSenEiKz5MSsEgzfR4k -XPOST -H "Content-Type: application/json" http://172.24.198.24:9200/index_media/_search?pretty -d '{"query":{"bool":{"must":[{"range":{"createtime":{"gt":"2020-01-01 00:00:00"}}}]}},"from":0,"size":10}'
