package sql

import (
	"bytes"
	"fmt"
	"github.com/yesAnd92/lwe/templates"
	"github.com/yesAnd92/lwe/utils"
	"go/format"
	"os"
	path2 "path"
	"path/filepath"
	"text/template"
)

type GoStructRenderData struct {
	*BaseParseDDL
	goStructTpl *template.Template
}

func NewGoStructRenderData() *GoStructRenderData {
	//加载实体对应的模板
	goStructTpl := templates.InitGoStructTpl()
	return &GoStructRenderData{
		goStructTpl: goStructTpl,
	}
}

func (g *GoStructRenderData) CovertSyntax(objInfos []*ObjInfo) error {
	for _, objInfo := range objInfos {
		for _, f := range objInfo.FieldInfos {
			//sql类型映射成java类型
			f.FieldType = utils.SqlToGoType(f.ColumnType)
			//sql字段名对应的Bean名字
			f.FieldName = utils.UderscoreToLowerCamelCase(f.ColumnName)
		}
	}
	return nil
}

func (g *GoStructRenderData) RenderData(objInfos []*ObjInfo) error {

	utils.MkdirIfNotExist(GENERATE_DIR)

	//生成的多个结构体先放到buffer中，最后一起写入文件
	bf := &bytes.Buffer{}

	//追加package import信息
	bf.Write([]byte(GO_TPL_HEAD))
	for _, objInfo := range objInfos {
		if err := g.goStructTpl.Execute(bf, objInfo); err != nil {
			return fmt.Errorf("execute template failed: %w", err)
		}
	}

	//使用objName作为生成的文件名
	fileName := path2.Join(GENERATE_DIR, GENERATE_GO_FILENAME)
	path, err := filepath.Abs(fileName)
	if err != nil {
		return fmt.Errorf("get abs path failed: %w", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create go file failed: %w", err)
	}
	defer f.Close()

	//按照go的风格进行格式化
	fmtBfBytes, err := format.Source(bf.Bytes())
	if err != nil {
		return fmt.Errorf("format source failed: %w", err)
	}

	if _, err := f.Write(fmtBfBytes); err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}
	return nil
}
