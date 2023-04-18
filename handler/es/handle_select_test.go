package es

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"reflect"
	"testing"
)

func TestHandleSelect(t *testing.T) {
	selectSqlMap := map[string]string{
		//单条件测试用例
		`select * from user order by age ,create_time desc  limit 10,10`:                                                        `{"query":{"bool":{"must":[{"match_all":{}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where last_name='yesand' order by age ,create_time desc  limit 10,10`:                               `{"query":{"bool":{"must":[{"match_phrase":{"last_name":"yesand"}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age>18 order by age ,create_time desc  limit 10,10`:                                           `{"query":{"bool":{"must":[{"range":{"age":{"gt":"18"}}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age<35 order by age ,create_time desc  limit 10,10`:                                           `{"query":{"bool":{"must":[{"range":{"age":{"lt":"35"}}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age>=18 order by age ,create_time desc  limit 10,10`:                                          `{"query":{"bool":{"must":[{"range":{"age":{"gte":"18"}}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age<=35 order by age ,create_time desc  limit 10,10`:                                          `{"query":{"bool":{"must":[{"range":{"age":{"lte":"35"}}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age != 35 order by age ,create_time desc  limit 10,10`:                                        `{"query":{"bool":{"must":[{"bool":{"must_not":[{"match_phrase":{"age":"35"}}]}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age in (18,19,20) order by age ,create_time desc  limit 10,10`:                                `{"query":{"bool":{"must":[{"terms":{"age":[18, 19, 20]}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age not in (18,19,20) order by age ,create_time desc  limit 10,10`:                            `{"query":{"bool":{"must":[{"bool":{"must_not":{"terms":{"age":[18, 19, 20]}}}}]}},"from":10,"size":10,"sort":[{"age":"asc"},{"create_time":"desc"}]}`,
		`select * from user where age not in (18,19,20) and name in ("y","z")`:                                                  `{"query":{"bool":{"must":[{"bool":{"must_not":{"terms":{"age":[18, 19, 20]}}}},{"terms":{"name":["y", "z"]}}]}},"from":0,"size":1}`,
		`select * from user where last_name like "yes" limit 10,10`:                                                             `{"query":{"bool":{"must":[{"match_phrase":{"last_name":"yes"}}]}},"from":10,"size":10}`,
		`select * from user where last_name not like "yes" limit 10,10`:                                                         `{"query":{"bool":{"must":[{"bool":{"must_not":{"match_phrase":{"last_name":"yes"}}}}]}},"from":10,"size":10}`,
		`select * from user where age between 18 and 35`:                                                                        `{"query":{"bool":{"must":[{"range":{"age":{"from":"18","to":"35"}}}]}},"from":0,"size":1}`,
		`select * from user where age between 18 and 35 and createTime between '2023-04-01 00:00:00' and '2023-04-10 00:00:00'`: `{"query":{"bool":{"must":[{"range":{"age":{"from":"18","to":"35"}}},{"range":{"createTime":{"from":""2023-04-01 00:00:00"","to":""2023-04-10 00:00:00""}}}]}},"from":0,"size":1}`,
		//多条件测试用例
		`select * from user where age>=18 and gender=1 and last_name = "yesAnd"`:   `{"query":{"bool":{"must":[{"range":{"age":{"gte":"18"}}},{"match_phrase":{"gender":"1"}},{"match_phrase":{"last_name":"yesAnd"}}]}},"from":0,"size":1}`,
		`select * from user where age>=18 and gender=1 or last_name = "yesAnd"`:    `{"query":{"bool":{"should":[{"bool":{"must":[{"range":{"age":{"gte":"18"}}},{"match_phrase":{"gender":"1"}}]}},{"match_phrase":{"last_name":"yesAnd"}}]}},"from":0,"size":1}`,
		`select * from user where age>=18 and ( gender=1 or last_name = "yesAnd")`: `{"query":{"bool":{"must":[{"range":{"age":{"gte":"18"}}},{"bool":{"should":[{"match_phrase":{"gender":"1"}},{"match_phrase":{"last_name":"yesAnd"}}]}}]}},"from":0,"size":1}`,
	}

	for sql, expect := range selectSqlMap {

		sta, _ := sqlparser.Parse(sql)
		dsl, _, _ := HandleSelect(sta.(*sqlparser.Select))
		fmt.Println(dsl)
		if !reflect.DeepEqual(dsl, expect) {
			t.Error("the generated dsl is not equal to expected", sql)
		}
	}
}
