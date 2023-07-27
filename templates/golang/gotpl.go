package java

import (
	"github.com/yesAnd92/lwe/utils"
	"strings"
	"text/template"
)

/**
* 本想将模板文件单独出来，每次使用时进行加载
* 但是build成可执行文件后，就无法访问到项目里的文件
* 因此将模板定义成一个字符串解决这个问题
**/

//生成go结构体的模板定义
var goStructTpl = `
//{{.ObjName}} {{.ObjComment}}
type {{.ObjName}} struct {
    {{/* 渲染字段 */ -}}
    {{range .FieldInfos}}
	{{upperFirst .FieldName}} {{.FieldType}}  ~gorm:"{{ if eq $.PrimaryField  .ColumnName}}primary_key;{{ end}}" json:"{{ .FieldName}}"~   {{if .FieldComment}}//{{.FieldComment}}{{- end -}}
	{{end}}

}

`

func InitGoStructTpl() *template.Template {
	//定义模板时使用~代替了`，这里替换回来
	goStructTpl = strings.ReplaceAll(goStructTpl, "~", "`")
	tpl := template.Must(template.New("goStructTpl").Funcs(utils.TemplateFunc).Parse(goStructTpl))
	return tpl
}
