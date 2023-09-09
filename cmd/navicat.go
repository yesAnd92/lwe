package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yesAnd92/lwe/handler/navicat"
	"io"
	"os"
)

/**
* navicat命令相关功能
 */
var (
	navicatCmd = &cobra.Command{
		Use:     `ncx`,
		Short:   `Decrypt password of connection in .ncx file`,
		Long:    `The config exported from Navicat is encrypted,ncx command can decrypt it`,
		Example: `lwe ncx ncx-file-path `,
		Args:    cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			ncxFilePath := args[0]
			data, dataErr := getNcxData(ncxFilePath)
			if dataErr != nil {
				fmt.Println(dataErr)
			}
			ncx, parseErr := navicat.ParseNcx(data)
			if parseErr != nil {
				cobra.CheckErr(fmt.Errorf("parse ncx file error: %s", parseErr))
			}

			for _, conn := range ncx.Conns {

				fmt.Printf("-----------%s-----------\n", conn.ConnectionName)
				fmt.Printf("DB type:  %s\n"+
					"Connection host: %s\n"+
					"Connection port: %s\n"+
					"Connection username: %s\n"+
					"Connection password: %s\n\n", conn.ConnType, conn.Host, conn.Port, conn.UserName, conn.Password)
			}

		},
	}
)

// getNcxData 获取ncx数据
func getNcxData(ncxFilePath string) (data []byte, err error) {
	fi, err := os.Open(ncxFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't find ncx file: %s", ncxFilePath))
	}
	defer fi.Close()

	buffer := bytes.Buffer{}
	br := bufio.NewReader(fi)
	for {
		lineByte, _, e := br.ReadLine()
		buffer.Write(lineByte)
		if e == io.EOF {
			break
		}
	}
	return buffer.Bytes(), nil
}
