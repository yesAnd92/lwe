package sql

import (
	"bytes"
	"fmt"
	"log"
	golang "lwe/templates/golang"
	"lwe/utils"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type GoStructRenderData struct {
	AbstractParseDDL
	goStructTpl *template.Template
}

func NewGoStructRenderData() *GoStructRenderData {
	//加载实体对应的模板
	goStructTpl := golang.InitGoStructTpl()
	return &GoStructRenderData{
		goStructTpl: goStructTpl,
	}
}

func (g *GoStructRenderData) CovertSyntax(objInfos []*ObjInfo) {
	for _, objInfo := range objInfos {
		for _, f := range objInfo.FieldInfos {
			//sql类型映射成java类型
			f.FieldType = utils.SqlToGoType(f.ColumnType)
			//sql字段名对应的Bean名字
			f.FieldName = utils.UderscoreToLowerCamelCase(f.ColumnName)
		}
	}
}

func (g *GoStructRenderData) RenderData(objInfos []*ObjInfo) {
	_, e := os.Stat(GENERATE_DIR)
	if os.IsNotExist(e) {
		//不存在，则新建一个目录
		err := os.Mkdir(GENERATE_DIR, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return
		}
	}

	//生成的多个结构体先放到buffer中，最后一起写入文件
	bf := &bytes.Buffer{}

	//追加package import信息
	bf.Write([]byte(GO_TPL_HEAD))
	for _, objInfo := range objInfos {
		g.goStructTpl.Execute(bf, objInfo)
	}

	//使用objName作为生成的文件名
	fileName := GENERATE_GO_FILENAME
	path, _ := filepath.Abs(fileName)
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Println("Create go file err", err)
		return
	}
	f.Write(bf.Bytes())

	//使用go提供的fmt命令对生成的文件进行格式化
	cmd := exec.Command("go", "fmt", fileName)
	cmd.Run()
}
