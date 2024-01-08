package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

//相对路径转换成绝对路径进行处理

func ToAbsPath(path string) string {

	if !filepath.IsAbs(path) {
		absDir, err := filepath.Abs(path)
		if err != nil {
			cobra.CheckErr(err)
		}
		path = absDir
	}
	return path
}

func MkdirIfNotExist(dir string) {

	if _, e := os.Stat(dir); os.IsNotExist(e) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("mkdir failed![%v]\n", err))
		}
	}
}
