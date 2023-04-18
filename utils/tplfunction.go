package utils

import (
	"strings"
	"text/template"
)

// TemplateFunc 自定义的函数，可以在模板渲染时使用
// 增加了模板的灵活度
var TemplateFunc = template.FuncMap{
	// 将第一个字母大写
	"upperFirst": strings.Title,
}
