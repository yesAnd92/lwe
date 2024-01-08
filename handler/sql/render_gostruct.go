package sql

import (
	"bytes"
	"github.com/yesAnd92/lwe/templates"
	"github.com/yesAnd92/lwe/utils"
	"go/format"
	"log"
	"os"
	path2 "path"
	"path/filepath"
	"text/template"
)

type GoStructRenderData struct {
	AbstractParseDDL
	goStructTpl *template.Template
}

func NewGoStructRenderData() *GoStructRenderData {
	//加载实体对应的模板
	goStructTpl := templates.InitGoStructTpl()
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

	utils.MkdirIfNotExist(GENERATE_DIR)

	//生成的多个结构体先放到buffer中，最后一起写入文件
	bf := &bytes.Buffer{}

	//追加package import信息
	bf.Write([]byte(GO_TPL_HEAD))
	for _, objInfo := range objInfos {
		g.goStructTpl.Execute(bf, objInfo)
	}

	//使用objName作为生成的文件名
	fileName := path2.Join(GENERATE_DIR, GENERATE_GO_FILENAME)
	path, _ := filepath.Abs(fileName)
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		log.Println("Create go file err", err)
		return
	}
	//按照go的风格进行格式化
	fmtBfBytes, _ := format.Source(bf.Bytes())

	f.Write(fmtBfBytes)
}
