package sql

import (
	"fmt"
	"github.com/yesAnd92/lwe/templates"
	"github.com/yesAnd92/lwe/utils"
	"log"
	"os"
	path2 "path"
	"path/filepath"
	"text/template"
)

type JavaRenderData struct {
	AbstractParseDDL
	JavaTpl *template.Template
}

func NewJavaRenderData() *JavaRenderData {
	//加载实体对应的模板
	javaTpl := templates.InitJavaTpl()
	return &JavaRenderData{
		JavaTpl: javaTpl,
	}
}

func (m *JavaRenderData) CovertSyntax(objInfos []*ObjInfo) {
	for _, objInfo := range objInfos {
		//sql类型映射成java类型
		//sql字段名对应的Bean名字
		for _, f := range objInfo.FieldInfos {
			f.FieldType = utils.SqlToJavaType(f.ColumnType)
			f.FieldName = utils.UderscoreToLowerCamelCase(f.ColumnName)
		}
	}
}

func (m *JavaRenderData) RenderData(objInfos []*ObjInfo) {
	utils.MkdirIfNotExist(GENERATE_DIR)

	for _, objInfo := range objInfos {
		//使用objName作为生成的文件名
		fileName := fmt.Sprintf(path2.Join(GENERATE_DIR, GENERATE_JAVA_FILENAME), objInfo.ObjName)
		path, _ := filepath.Abs(fileName)
		f, err := os.Create(path)
		defer f.Close()

		if err != nil {
			log.Println("Create java file err", err)
			return
		}
		m.JavaTpl.Execute(f, objInfo)
	}
}
