package utils

//sql->java
func mysqlToJavaTypeMap() map[string]string {
	mysqlJavaTypeMap := map[string]string{
		"bigint":    "Long",
		"int":       "Integer",
		"tinyint":   "Integer",
		"smallint":  "Integer",
		"mediumint": "Integer",
		"integer":   "Integer",
		//小数
		"float":   "Float",
		"double":  "Double",
		"decimal": "Double",
		//bool
		"bit": "Boolean",
		//字符串
		"char":       "String",
		"varchar":    "String",
		"tinytext":   "String",
		"text":       "String",
		"mediumtext": "String",
		"longtext":   "String",
		//日期
		"date":      "Date",
		"datetime":  "Date",
		"timestamp": "Date"}
	return mysqlJavaTypeMap
}

//sql->go
func mysqlToGoTypeMap() map[string]string {
	mysqlJavaTypeMap := map[string]string{
		"bigint":    "int",
		"int":       "int",
		"tinyint":   "int64",
		"smallint":  "uint",
		"mediumint": "uint",
		"integer":   "uint",
		//小数
		"float":   "float64",
		"double":  "float64",
		"decimal": "string",
		//bool
		"bit": "bool",
		//字符串
		"varbinary":  "string", //二进制流
		"char":       "string",
		"varchar":    "string",
		"tinytext":   "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		//日期
		"date":      "time.Time",
		"datetime":  "time.Time",
		"timestamp": "time.Time"}
	return mysqlJavaTypeMap
}

func SqlToJavaType(sqlType string) string {
	return mysqlToJavaTypeMap()[sqlType]
}

func SqlToGoType(sqlType string) string {
	return mysqlToGoTypeMap()[sqlType]
}
