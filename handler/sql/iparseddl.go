package sql

// IParseDDL 解析生成目标文件的核心接口
type IParseDDL interface {

	// ParseDDL 解析DDL文本
	//args 命令行传入的参数，如：注释中的author字段
	ParseDDL(sqlText string, args map[string]interface{}) (*ObjInfo, error)

	// CovertSyntax 转换到不同语言的字段类型
	// 比如sql中int对应到Java中的Integer，对应到go中的int32等
	CovertSyntax(info []*ObjInfo) error

	// RenderData 根据模版渲染数据
	RenderData(info []*ObjInfo) error
}

// 定义sql所用的常量
const (
	//生成文件所在的目录
	GENERATE_DIR = `./lwe-generate-file`
	//生成JAVA文件名
	GENERATE_JAVA_FILENAME = "%s.java"
	//生成Go文件名
	GENERATE_GO_FILENAME = "lwe_struct.go"
	//go模板头内容
	GO_TPL_HEAD = `package lwe  
					import ("time")`
)

// ObjInfo 表映射的对象及其字段信息
type ObjInfo struct {
	TableName    string                 //  表名
	ObjName      string                 //	 对象名
	ObjComment   string                 //	 对象的注释
	GenerateDate string                 //  生成日期
	PrimaryField string                 //  主键对应的字段
	FieldInfos   []*FieldInfo           //  字段的切片
	Args         map[string]interface{} //  命令行传入的参数
}

// FieldInfo 字段信息
type FieldInfo struct {
	ColumnName   string //  列名
	ColumnType   string //  列类型
	FieldName    string //  字段名
	FieldType    string //  字段类型
	FieldComment string //  字段的注释
}
