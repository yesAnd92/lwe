package templates

import (
	_ "embed"
	"github.com/yesAnd92/lwe/utils"
	"text/template"
)

//go:embed format_go.tpl
var goStructTpl string

//go:embed format_java.tpl
var javaStructTpl string

// InitGoStructTpl go template
func InitGoStructTpl() *template.Template {
	tpl := template.Must(template.New("goTpl").Funcs(utils.TemplateFunc).Parse(goStructTpl))
	return tpl
}

// InitJavaTpl java template
func InitJavaTpl() *template.Template {
	tpl := template.Must(template.New("javaTpl").Funcs(utils.TemplateFunc).Parse(javaStructTpl))
	return tpl
}
