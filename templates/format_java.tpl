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