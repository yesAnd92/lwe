package sync

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/yesAnd92/lwe/utils"
	"os"
	"strings"
	"sync"
)

// struct{}不占用空间
type void struct{}

type CompareThenDoIfa interface {
	Do(*Fsync)
}

type Fsync struct {
	sourceDir string
	targetDir string
	commonSet map[string]void
	sUnique   *[]string
	tUnique   *[]string
}

func InitFsync(sourceDir, targetDir string) *Fsync {
	return &Fsync{
		sourceDir: utils.ToAbsPath(sourceDir),
		targetDir: utils.ToAbsPath(targetDir),
		commonSet: make(map[string]void),
		sUnique:   &[]string{},
		tUnique:   &[]string{},
	}
}

func (f *Fsync) DiffDir() {
	sourceDir := f.sourceDir
	targetDir := f.targetDir

	wg := sync.WaitGroup{}
	wg.Add(2)
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

	if len(sourceFileSet) < len(targetFileSet) {
		for s := range sourceFileSet {
			rlvSource := strings.TrimPrefix(s, sourceDir)
			if _, ok := targetFileSet[targetDir+rlvSource]; ok {
				//common
				f.commonSet[rlvSource] = void{}
			}
		}
	} else {
		for t := range targetFileSet {
			rlvTarget := strings.TrimPrefix(t, targetDir)
			if _, ok := sourceFileSet[sourceDir+rlvTarget]; ok {
				//common
				f.commonSet[rlvTarget] = void{}
			}
		}
	}

	for s := range sourceFileSet {
		sp := strings.TrimPrefix(s, sourceDir)
		if _, ok := f.commonSet[sp]; !ok {
			*f.sUnique = append(*f.sUnique, sp)
		}
	}

	for t := range targetFileSet {
		tp := strings.TrimPrefix(t, targetDir)
		if _, ok := f.commonSet[tp]; !ok {
			*f.tUnique = append(*f.tUnique, tp)
		}
	}
}

func (f *Fsync) Sync(thenDo CompareThenDoIfa) {
	thenDo.Do(f)
}

type DisplayCompareThenDo struct {
}

func (d *DisplayCompareThenDo) Do(fsync *Fsync) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Source", " VS", "Target"})

	sourceDir := fsync.sourceDir
	targetDir := fsync.targetDir

	for k, _ := range fsync.commonSet {
		t.AppendRow(table.Row{sourceDir + k, "<==>", targetDir + k})
	}

	for _, v := range *fsync.sUnique {
		t.AppendRow(table.Row{sourceDir + v, "===>", ""})
	}

	for _, v := range *fsync.tUnique {
		t.AppendRow(table.Row{"", "<===", targetDir + v})
	}

	if t.Length() > 0 {
		t.Render()
	} else {
		fmt.Println("Source and target dir are empty dir!")
	}
	fmt.Println()
}

type CopyCompareThenDo struct {
}

func (c *CopyCompareThenDo) Do(fsync *Fsync) {

	sourceDir := fsync.sourceDir
	targetDir := fsync.targetDir

	//source unique to target
	for _, path := range *fsync.sUnique {
		written, err := utils.Copy(sourceDir+path, targetDir+path)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			fmt.Printf("%s\tcopy finished!,total:%2fKB\n", targetDir+path, float64(written)/1014)
		}
	}
	fmt.Println("All file copy finished!")

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
		}

	}
}
