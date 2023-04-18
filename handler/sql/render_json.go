package sql

import (
	"encoding/json"
	"fmt"
	"lwe/utils"
	"time"
)

type JsonRenderData struct {
	AbstractParseDDL
}

func NewJsonRenderData() *JsonRenderData {
	return &JsonRenderData{}
}

func (j *JsonRenderData) RenderData(objInfos []*ObjInfo) {
	for _, objInfo := range objInfos {
		infos := objInfo.FieldInfos
		filedMap := make(map[string]interface{}, len(infos))
		for _, info := range infos {
			switch info.FieldType {

			//塞入默认值
			case "Integer":
				filedMap[info.FieldName] = 0
			case "Long":
				filedMap[info.FieldName] = 0
			case "Float":
				filedMap[info.FieldName] = 1.0
			case "Double":
				filedMap[info.FieldName] = 1.0
			case "String":
				filedMap[info.FieldName] = info.FieldComment
			case "Date":
				filedMap[info.FieldName] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
		marshal, _ := json.MarshalIndent(filedMap, "", "  ")
		fmt.Println(string(marshal))
	}
}
func (j *JsonRenderData) CovertSyntax(objInfos []*ObjInfo) {

	for _, objInfo := range objInfos {
		//sql类型映射成java类型
		//sql字段名对应的Bean名字
		for _, f := range objInfo.FieldInfos {
			f.FieldType = utils.SqlToJavaType(f.ColumnType)
			f.FieldName = utils.UderscoreToLowerCamelCase(f.ColumnName)
		}
	}
}
