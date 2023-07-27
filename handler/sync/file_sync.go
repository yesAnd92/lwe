package sync

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findAllFile(dir string) []string {

	filePaths := &[]string{}
	// 打开目录并获取文件信息
	walk(dir, filePaths)

	for _, path := range *filePaths {
		fmt.Println(path)
		fmt.Println(strings.Trim(path, dir))

	}

	return nil
}

func walk(dir string, filePaths *[]string) {
	// 打开目录并获取文件信息
	finfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range finfo {

		//遇到目录递归遍历
		if info.IsDir() {
			walk(dir+"/"+info.Name(), filePaths)
		} else {
			*filePaths = append(*filePaths, dir+"/"+info.Name())
			//fmt.Println(dir + "/" + info.Name())
		}

	}
}
