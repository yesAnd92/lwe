package es

import (
	"errors"
	"fmt"
	"github.com/xwb1989/sqlparser"
	"strings"
)

func HandleSelect(sel *sqlparser.Select) (dsl string, esType string, err error) {

	var defaultQueryDsl = `{"bool":{"must":[{"match_all":{}}]}}`
	var queryDsl string

	//目的是告诉下一层级这是个root节点
	var root sqlparser.Expr
	//处理where子句，构造queryDsl
	if sel.Where != nil {
		queryDsl, err = handlerWhere(&sel.Where.Expr, true, &root)
		if err != nil {
			return "", "", err
		}
	}

	if queryDsl == "" {
		queryDsl = defaultQueryDsl
	}

	//处理表名
	if len(sel.From) > 1 {
		return "", "", errors.New("not support multiple table now")
	}
	esType = sqlparser.String(sel.From)

	//处理分页参数
	offset, rowCount := "0", "1"
	if sel.Limit != nil {
		offset = sqlparser.String(sel.Limit.Offset)
		rowCount = sqlparser.String(sel.Limit.Rowcount)
	}

	//处理排序,可能有多个字段
	var orderByArr []string
	if sel.OrderBy != nil {
		for _, orderExpr := range sel.OrderBy {
			//取出每对<排序字段,升降序>进行拼接
			colName := sqlparser.String(orderExpr.Expr)
			direction := orderExpr.Direction
			orderByArr = append(orderByArr, fmt.Sprintf(`{"%v":"%v"}`, colName, direction))
		}
	}

	var resultMap = map[string]string{
		"query": queryDsl,
		"from":  offset,
		"size":  rowCount,
	}
	if len(orderByArr) > 0 {
		//如果存在排序字段，多个排序条件是个数组，用"[]"包起来
		querySort := fmt.Sprintf("[%v]", strings.Join(orderByArr, ","))
		resultMap["sort"] = querySort
	}

	//指定关键字的顺序来保证生成结果的稳定
	//稳定的顺序可以使风格更统一，也便于写测试用例
	var desiredSequence = []string{"query", "from", "size", "sort", "aggregations"}
	var resultArr []string
	for _, key := range desiredSequence {
		if value, ok := resultMap[key]; ok {
			resultArr = append(resultArr, fmt.Sprintf(`"%v":%v`, key, value))
		}
	}
	dsl = fmt.Sprintf("{%v}", strings.Join(resultArr, ","))
	return dsl, esType, nil
}

func handlerWhere(expr *sqlparser.Expr, isTop bool, parent *sqlparser.Expr) (string, error) {
	if expr == nil {
		return "", errors.New("handlerWhere: expr can't be nil here ")
	}
	switch (*expr).(type) {

	//处理表达式
	case *sqlparser.ComparisonExpr:
		return handleSelectWhereComparisonExpr(expr, isTop)
	case *sqlparser.AndExpr:
		return handleSelectWhereAndExpr(expr, isTop, parent)
	case *sqlparser.OrExpr:
		return handleSelectWhereOrExpr(expr, isTop, parent)
	case *sqlparser.RangeCond:
		return handleSelectWhereRangeExpr(expr, isTop, parent)
	case *sqlparser.ParenExpr:
		return handleSelectWhereParentExpr(expr, isTop, parent)
	default:
		return "", errors.New(fmt.Sprintf(`[%v] is supported currently! `, sqlparser.String(*expr)))
	}

	return "", nil

}

//处理between范围，例如between a and b
func handleSelectWhereRangeExpr(expr *sqlparser.Expr, isTop bool, parent *sqlparser.Expr) (string, error) {

	rangeCondExpr := (*expr).(*sqlparser.RangeCond)

	var colNameStr, fromStr, toStr string
	if colName, ok := rangeCondExpr.Left.(*sqlparser.ColName); ok {
		colNameStr = sqlparser.String(colName)
	} else {
		return "", errors.New("invalidated comparison expression")
	}

	fromStr = sqlparser.String(rangeCondExpr.From)
	fromStr = strings.NewReplacer("'", "\"").Replace(fromStr)

	toStr = sqlparser.String(rangeCondExpr.To)
	toStr = strings.NewReplacer("'", "\"").Replace(toStr)

	resultStr := fmt.Sprintf(`{"range":{"%v":{"from":"%v","to":"%v"}}}`, colNameStr, fromStr, toStr)

	if isTop {
		resultStr = fmt.Sprintf(`{"bool":{"must":[%v]}}`, resultStr)
	}

	return resultStr, nil

}

//处理带括号的语句，比如 where a=1 and (b=1 or c=1 )
func handleSelectWhereParentExpr(expr *sqlparser.Expr, isTop bool, parent *sqlparser.Expr) (string, error) {
	parenExpr := (*expr).(*sqlparser.ParenExpr)
	boolExpr := parenExpr.Expr
	return handlerWhere(&boolExpr, isTop, parent)

}

func handleSelectWhereAndExpr(expr *sqlparser.Expr, isTop bool, parent *sqlparser.Expr) (string, error) {
	andExpr := (*expr).(*sqlparser.AndExpr)
	var leftStr, rightStr string
	var err error

	//AndExpr的左右两边可能是单独的ComparisonExpr，也可能还是个AndExpr等，
	//因此还可以调用handlerWhere，即递归

	//处理and左边的表达式
	leftStr, err = handlerWhere(&andExpr.Left, false, expr)
	if err != nil {
		return "", err
	}

	//处理and右边的表达式
	rightStr, err = handlerWhere(&andExpr.Right, false, expr)
	if err != nil {
		return "", err
	}

	//拼接
	var resultStr string
	if leftStr == "" || rightStr == "" {
		resultStr = leftStr + rightStr
	} else {
		resultStr = leftStr + "," + rightStr
	}

	//如果上一级也是and关系，则没必要在拼接bool must
	if _, ok := (*parent).(*sqlparser.AndExpr); ok {
		return resultStr, nil
	}

	return fmt.Sprintf(`{"bool":{"must":[%v]}}`, resultStr), nil
}

func handleSelectWhereOrExpr(expr *sqlparser.Expr, isTop bool, parent *sqlparser.Expr) (string, error) {
	orExpr := (*expr).(*sqlparser.OrExpr)
	var leftStr, rightStr string
	var err error

	//AndExpr的左右两边可能是单独的ComparisonExpr，也可能还是个AndExpr等，
	//因此还可以调用handlerWhere，即递归

	//处理and左边的表达式
	leftStr, err = handlerWhere(&orExpr.Left, false, expr)
	//处理and右边的表达式
	rightStr, err = handlerWhere(&orExpr.Right, false, expr)

	//拼接
	var resultStr string
	if leftStr == "" || rightStr == "" {
		resultStr = leftStr + rightStr
	} else {
		resultStr = leftStr + "," + rightStr
	}

	//如果上一级也是or关系，则没必要在拼接bool must
	if _, ok := (*parent).(*sqlparser.OrExpr); ok {
		return resultStr, nil
	}

	return fmt.Sprintf(`{"bool":{"should":[%v]}}`, resultStr), err
}

func handleSelectWhereComparisonExpr(expr *sqlparser.Expr, isTop bool) (string, error) {
	comparisonExpr := (*expr).(*sqlparser.ComparisonExpr)
	var colNameStr string
	var valueStr string
	if colName, ok := comparisonExpr.Left.(*sqlparser.ColName); ok {
		colNameStr = sqlparser.String(colName)
	} else {
		return "", errors.New("invalidated comparison expression")
	}

	//由于操作符右边的可能性比较多，单独进行处理
	valueStr, err := buildComparisonExprRight(comparisonExpr.Right)
	if err != nil {
		return "", err
	}

	var comparisonStr string
	//根据不同的比较符号，返回拼接对应的语句
	switch comparisonExpr.Operator {

	case "=":
		//等号对应到dsl的match_phrase
		//{"match_phrase":{k:v}} 和 {"match_phrase":{k:{"query":v}}}等价，这里都是用第一种写法
		//match_phrase查询首先解析查询字符串来产生一个词条列表。然后会搜索所有的词条
		//但只保留含有了所有搜索词条的文档，并且词条的位置要邻接
		comparisonStr = fmt.Sprintf(`{"match_phrase":{"%v":"%v"}}`, colNameStr, valueStr)

	case ">":
		//范围类操作符对应到dsl的range
		comparisonStr = fmt.Sprintf(`{"range":{"%v":{"gt":"%v"}}}`, colNameStr, valueStr)
	case "<":
		comparisonStr = fmt.Sprintf(`{"range":{"%v":{"lt":"%v"}}}`, colNameStr, valueStr)

	case ">=":
		//范围类操作符对应到dsl的range
		comparisonStr = fmt.Sprintf(`{"range":{"%v":{"gte":"%v"}}}`, colNameStr, valueStr)
	case "<=":
		comparisonStr = fmt.Sprintf(`{"range":{"%v":{"lte":"%v"}}}`, colNameStr, valueStr)

	case "!=":
		//!=对应到dsl的must_not
		comparisonStr = fmt.Sprintf(`{"bool":{"must_not":[{"match_phrase":{"%v":"%v"}}]}}`, colNameStr, valueStr)

	case "in":
		//in对应到dsl的terms
		comparisonStr = fmt.Sprintf(`{"terms":{"%v":[%v]}}`, colNameStr, valueStr)

	case "not in":
		//in对应到dsl的terms
		comparisonStr = fmt.Sprintf(`{"bool":{"must_not":{"terms":{"%v":[%v]}}}}`, colNameStr, valueStr)

	case "like":
		//like对应到dsl的match_phrase
		comparisonStr = fmt.Sprintf(`{"match_phrase":{"%v":"%v"}}`, colNameStr, valueStr)
	case "not like":
		//like对应到dsl的match_phrase
		comparisonStr = fmt.Sprintf(`{"bool":{"must_not":{"match_phrase":{"%v":"%v"}}}}`, colNameStr, valueStr)
	}

	if isTop {
		comparisonStr = fmt.Sprintf(`{"bool":{"must":[%v]}}`, comparisonStr)
	}
	return comparisonStr, nil
}

func buildComparisonExprRight(expr sqlparser.Expr) (string, error) {

	var valueStr string

	switch expr.(type) {

	case *sqlparser.SQLVal:
		valueStr = sqlparser.String(expr)
		//统一去掉引号类字符
		valueStr = strings.NewReplacer("`", "", "'", "", "\"", "").Replace(valueStr)

	case sqlparser.ValTuple:
		//sqlparser.ValTuple为啥不是指针了？？
		valueStr = sqlparser.String(expr)
		//去掉首尾的括号、引号类字符
		valueStr = strings.NewReplacer("(", "", ")", "", "'", "\"").Replace(valueStr)
	}

	return valueStr, nil
}
