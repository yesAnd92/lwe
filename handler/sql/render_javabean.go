package sql

import (
	"fmt"
	"github.com/yesAnd92/lwe/templates/java"
	"github.com/yesAnd92/lwe/utils"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type JavaRenderData struct {
	AbstractParseDDL
	JavaTpl *template.Template
}

func NewJavaRenderData() *JavaRenderData {
	//加载实体对应的模板
	javaTpl := java.InitJavaTpl()
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
	_, e := os.Stat(GENERATE_DIR)
	if os.IsNotExist(e) {
		//不存在，则新建一个目录
		err := os.Mkdir(GENERATE_DIR, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return
		}
	}

	for _, objInfo := range objInfos {
		//使用objName作为生成的文件名
		fileName := fmt.Sprintf(GENERATE_JAVA_FILENAME, objInfo.ObjName)
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
