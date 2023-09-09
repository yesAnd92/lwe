//{{.ObjName}} {{.ObjComment}}
type {{.ObjName}} struct {
    {{/* 渲染字段 */ -}}
    {{range .FieldInfos}}
	{{upperFirst .FieldName}} {{.FieldType}}  `gorm:"{{ if eq $.PrimaryField  .ColumnName}}primary_key;{{ end}}" json:"{{ .FieldName}}"`   {{if .FieldComment}}//{{.FieldComment}}{{- end -}}
	{{end}}

}
