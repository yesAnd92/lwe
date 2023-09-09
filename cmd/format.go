package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/sql"
	"io"
	"os"
	"strings"
)

/**
* sql命令相关功能
 */

var (
	target string
	author string
	sqlCmd = &cobra.Command{
		Use:     `fmt`,
		Short:   `Generate the specified file based on SQL`,
		Long:    `Generate the specified file based on SQL. Such as Java Entity,Go struct and so on`,
		Example: `lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]`,
		Args:    cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			sqlFilePath := args[0]
			sqlCtxList, err := DataCleansing(sqlFilePath)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("can't find SQL file: %s", sqlFilePath))
			}

			//由target参数找到对应的handle
			handle, err := GetParser(target)
			if err != nil {
				cobra.CheckErr(err)
			}

			//设置作者参数
			params := make(map[string]interface{})
			if len(author) > 0 {
				params["author"] = author
			}

			sql.DoParse(handle, sqlCtxList, params)
			fmt.Println("File generated successfully!")
		},
	}
)

func GetParser(target string) (sql.IParseDDL, error) {
	var handle sql.IParseDDL
	switch target {
	case "java":
		handle = sql.NewJavaRenderData()
	case "go":
		handle = sql.NewGoStructRenderData()
	case "json":
		handle = sql.NewJsonRenderData()
	}
	if handle == nil {
		return nil, errors.New("target " + target + " param error!")
	}
	return handle, nil
}

// DataCleansing 清洗数据，摘取create语句
func DataCleansing(sqlFilePath string) (ctxList []string, err error) {
	fi, err := os.Open(sqlFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't find SQL file: %s", sqlFilePath))
	}
	defer fi.Close()

	//存放所有的创建语句
	var sqlCtxArr []string
	//创建语句
	var createStr string
	br := bufio.NewReader(fi)
	for {
		lineByte, _, e := br.ReadLine()
		//找到create语句的开始行，一直读取到这个语句结束
		if strings.HasPrefix(string(lineByte), "CREATE") || strings.HasPrefix(string(lineByte), "create") {
			createStr += string(lineByte)
			for {
				lb, _, e2 := br.ReadLine()
				//本条create语句结束，放到结果容器中，继续寻找下一个create语句
				if len(lb) == 0 || e2 == io.EOF {
					sqlCtxArr = append(sqlCtxArr, createStr)
					createStr = ""
					break
				}
				createStr += string(lb)
			}
		}
		if e == io.EOF {
			break
		}
	}
	return sqlCtxArr, nil
}

func init() {

	sqlCmd.PersistentFlags().StringVarP(&target, "target", "t", "java", "The type[java|json|go] of generate the sql")
	sqlCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "Comment for author information will be added to the generated file")
}
