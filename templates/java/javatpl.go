package java

import (
	"github.com/yesAnd92/lwe/utils"
	"text/template"
)

/**
* 本想将模板文件单独出来，每次使用时进行加载
* 但是build成可执行文件后，就无法访问到项目里的文件
* 因此将模板定义成一个字符串解决这个问题
**/

//生成java实体的模板定义
var javaTpl = `
import java.util.Date;
import java.util.List;
import java.io.Serializable;
import javax.persistence.Column;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.GeneratedValue;

/**
 * @Description {{.ObjComment}}
 * @Author {{if .Args.author}} {{.Args.author}} {{- end}}
 * @Date  {{.GenerateDate}}
 */
@Table ( name ="{{.TableName}}" )
public class {{.ObjName}} implements Serializable {

    private static final long serialVersionUID = 1L;

    {{/* 渲染字段 */ -}}
    {{range .FieldInfos}}
	{{ if eq $.PrimaryField  .ColumnName}}@Id {{ end}}
	@Column(name = "{{.ColumnName}}" )
    private {{.FieldType}} {{.FieldName}};	{{if .FieldComment}}//{{.FieldComment}}{{- end -}}
    {{end}}

    {{/* 渲染字段对应的setter、getter */ -}}
    {{range .FieldInfos}}
    public {{ .FieldType}} get{{upperFirst .FieldName}}() {
        return {{.FieldName}};
    }

    public void set{{upperFirst .FieldName}}({{.FieldType}} {{.FieldName}}) {
        this.{{.FieldName}} = {{.FieldName}};
    }
    {{end}}
}
`

func InitJavaTpl() *template.Template {
	tpl := template.Must(template.New("javaTpl").Funcs(utils.TemplateFunc).Parse(javaTpl))
	return tpl
}
