package sql

import (
	"errors"
	"fmt"
	"github.com/yesAnd92/lwe/utils"
	"regexp"
	"strings"
	"time"
)

// AbstractParseDDL IParseDDL 接口抽象实现
type AbstractParseDDL struct {
}

// DoParse 定义了整个解析、生成的流程
func DoParse(parse IParseDDL, sqlTextArr []string, args map[string]interface{}) {
	//解析ddl文本
	var objInfos = make([]*ObjInfo, 0)
	for _, sqlText := range sqlTextArr {
		objInfo, err := parse.ParseDDL(sqlText, args)
		if err != nil {
			fmt.Println(err)
		}
		objInfos = append(objInfos, objInfo)
	}
	//适配不同的生成类型，由子类实现
	parse.CovertSyntax(objInfos)

	//渲染数据
	parse.RenderData(objInfos)
}

func (a AbstractParseDDL) ParseDDL(sqlText string, args map[string]interface{}) (*ObjInfo, error) {

	if sqlText == "" {
		return nil, errors.New("SQL不能为空")
	}

	//替换掉误输入字符、空格等
	r1 := strings.NewReplacer("'", "`", "，", ",", "\n", "", "\t", "")
	sqlText = strings.ToLower(r1.Replace(sqlText))

	//新增处理create table if not exists members情况
	sqlText = strings.ReplaceAll(sqlText, "if not exists", "")

	//表名
	var tableName string
	if strings.Contains(sqlText, "table") {
		tableName = sqlText[strings.Index(sqlText, "table")+5 : strings.Index(sqlText, "(")]
	}

	if strings.Contains(tableName, "`") {
		tableName = tableName[strings.Index(tableName, "`")+1 : strings.LastIndex(tableName, "`")]
	} else {
		//空格开头的，需要替换掉\n\t空格
		tableName = strings.NewReplacer("\n", "", "\t", "").Replace(tableName)
	}
	var originTableName = tableName

	//转化为类名
	className := utils.UderscoreToUpperCamelCase(tableName)

	//表注释
	var tableComment string
	//mysql是comment=,pgsql/oracle是comment on table
	if strings.Contains(sqlText, "comment=") {
		tableComment = sqlText[strings.Index(sqlText, "comment=")+8:]
	} else if strings.Contains(sqlText, "comment on table") {
		tableComment = sqlText[strings.Index(sqlText, "comment on table")+17:]
	} else {
		//没有表注释，使用表名作为缺省值
		tableComment = originTableName
	}
	tableComment = strings.NewReplacer("`", "", ";", "").Replace(tableComment)

	//主键对应的字段名
	var primaryFiled string
	//因为primary key ()中亦可能出现",",因此提前处理掉，防止被当成切割符号切割
	//using btree可有可无都允许
	//go标准regexp库中如果支持零宽断言，可以直接使用(?<=key \(`)\S*(?=`\))匹配出主键字段来
	primaryFiledRe := regexp.MustCompile("primary key \\((.*?)\\)( using btree)*( comment `(.*?)`)*")
	primaryFiledFound := primaryFiledRe.FindAllString(sqlText, -1)
	if len(primaryFiledFound) == 1 {
		filedStr := primaryFiledFound[0]
		primaryFiled = filedStr[strings.Index(filedStr, "(`")+2 : strings.Index(filedStr, "`)")]
		//移除sqlText中的关于primary key部分
		sqlText = strings.ReplaceAll(sqlText, filedStr, "")
	}

	//移除索引
	//普通索引，全文索引，唯一索引都转换成普通索引统一处理
	sqlText = strings.NewReplacer("unique key", "key", "fulltext key", "key").Replace(sqlText)
	//所有后面也可以有评论comment
	indexRe := regexp.MustCompile("key `(.*?)`\\)( using btree)*( comment `(.*?)`)*")
	indexFound := indexRe.FindAllString(sqlText, -1)
	for _, indexStr := range indexFound {
		//索引相关与生成文件无关，这里处理掉避免产生干扰
		sqlText = strings.NewReplacer(indexStr, "").Replace(sqlText)
	}

	//截取字段部分
	var filedListTmp = sqlText[strings.Index(sqlText, "(")+1 : strings.LastIndex(sqlText, ")")]

	// 匹配 comment，替换备注里的小逗号, 防止被当成切割符号切割
	re := regexp.MustCompile("comment `(.*?)\\`")
	found := re.FindAllString(filedListTmp, -1)
	for _, filedStr := range found {
		//使用中文逗号，替换注释中的英文逗号
		filedStrNew := strings.ReplaceAll(filedStr, ",", "，")
		//再替换回filedLisTmp中的相应位置
		filedListTmp = strings.ReplaceAll(filedListTmp, filedStr, filedStrNew)
	}

	//匹配 double(20,2) 这种精度描述，替换备注里的小逗号, 防止被当成切割符号切割
	re = regexp.MustCompile("\\([0-9]+,[0-9]*\\)")
	found = re.FindAllString(filedListTmp, -1)
	for _, filedStr := range found {
		//直接替换掉(x,x)
		filedListTmp = strings.ReplaceAll(filedListTmp, filedStr, "")
	}
	//解析、组装各个字段
	var fileds = make([]*FieldInfo, 0)
	filedArr := strings.Split(filedListTmp, ",")
	for _, commandline := range filedArr {

		//去除字段行定义语句的前后空格
		commandline = strings.TrimSpace(commandline)
		if len(commandline) == 0 {
			continue
		}

		//列名
		columnName := commandline[0:strings.Index(commandline, " ")]
		columnName = strings.ReplaceAll(columnName, "`", "")

		//columnType
		columnType := strings.Split(commandline, " ")[1]
		if strings.Contains(columnType, "(") {
			columnType = columnType[0:strings.Index(columnType, "(")]
		}

		//filedComment
		var filedComment string
		//mysql的字段注释位于行末
		if strings.Contains(commandline, "comment `") {
			filedComment = commandline[strings.Index(commandline, "comment")+7:]
			filedComment = strings.ReplaceAll(strings.TrimSpace(filedComment), "`", "")
		}

		fileds = append(fileds, &FieldInfo{
			ColumnName:   columnName,
			ColumnType:   columnType,
			FieldComment: filedComment,
		})
	}

	obj := &ObjInfo{
		TableName:    originTableName,
		ObjName:      className,
		ObjComment:   tableComment,
		GenerateDate: time.Now().Format("2006/01/02 15:04"),
		PrimaryField: primaryFiled,
		FieldInfos:   fileds,
		Args:         args,
	}

	return obj, nil
}

func (a AbstractParseDDL) CovertSyntax(info *ObjInfo) {
	//TODO implement me
	panic("implement me")
}

func (a AbstractParseDDL) RenderData(info *ObjInfo) {
	//TODO implement me
	panic("implement me")
}
