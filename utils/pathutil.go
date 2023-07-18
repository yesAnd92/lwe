package utils

import (
	"github.com/spf13/cobra"
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
