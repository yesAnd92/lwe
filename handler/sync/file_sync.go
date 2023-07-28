package sync

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// struct{}不占用空间
type void struct{}

func compareDir(sourceDir, targetDir string) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	//var sFilePath = &[]string{}
	//var tFilePath = &[]string{}
	var sourceFileSet = make(map[string]void)
	var targetFileSet = make(map[string]void)
	go func() {
		defer wg.Done()
		findAllFile(sourceDir, sourceFileSet)
	}()

	go func() {
		defer wg.Done()
		findAllFile(targetDir, targetFileSet)
	}()
	wg.Wait()

	fmt.Println(sourceFileSet)
	fmt.Println(targetFileSet)
	//收集targetDir中已经有的文件

	//var common = &[]string{}
	//var sUnique = &[]string{}
	//var tUnique = &[]string{}

	if len(sourceFileSet) < len(targetFileSet) {

	}

}

func findAllFile(dir string, re map[string]void) {

	// 打开目录并获取文件信息
	walk(dir, re)

}

func walk(dir string, pathMap map[string]void) {
	// 打开目录并获取文件信息
	dirEntry, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range dirEntry {
		//根据文件或文件夹名是否"."开头判断是否隐藏内容，隐藏在不同步
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		//遇到目录递归遍历
		if file.IsDir() {
			walk(dir+"/"+file.Name(), pathMap)
		} else {
			pathMap[dir+"/"+file.Name()] = void{}
			//*filePaths = append(*filePaths, dir+"/"+file.Name())
		}

	}
}

func copy() {
	// 创建源文件
	_, _ = os.Create("src.txt")
	// 打开源文件
	file1, err1 := os.Open("src.txt")
	if err1 != nil {
		fmt.Println(err1)
	}

	// 创建目标文件
	file2, err2 := os.OpenFile("dest.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	//使用结束关闭文件
	defer file1.Close()
	defer file2.Close()
	n, e := io.Copy(file2, file1)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("拷贝成功。。。，拷贝字节数：", n)
	}

}
