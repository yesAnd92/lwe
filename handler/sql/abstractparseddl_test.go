package sql

import (
	"testing"
)

func GetTestCase() []string {
	return []string{
		// TODO: Add test cases.
		`CREATE TABLE 'student_info' (
				  'id' int(11) NOT NULL AUTO_INCREMENT COMMENT '用户编号,学号',
				  'class_id' varchar(255) NOT NULL COMMENT '班级id',
				  'user_name' varchar(255) NOT NULL COMMENT '用户名',
				  'status' tinyint(1) NOT NULL COMMENT '状态',
				  'create_time' datetime NOT NULL COMMENT '创建时间',
				  PRIMARY KEY ('id') 
				) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学生信息'`,
		`CREATE TABLE 'class_info' (
				  'id' int(11) NOT NULL AUTO_INCREMENT COMMENT '用户编号,学号',
				  'name' varchar(255) NOT NULL COMMENT '班级名',
				  'create_time' datetime NOT NULL COMMENT '创建时间',
    			  'total_size' double(20,2) DEFAULT NULL,
				  PRIMARY KEY ('id') USING BTREE COMMENT '主键'
				) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='班级信息'`,
		//索引
		`CREATE TABLE 'template_role' (
				  'role_id' varchar(32) NOT NULL COMMENT '角色id',
				  'template_id' varchar(32) NOT NULL COMMENT '模板id',
				  'create_time' datetime DEFAULT NULL COMMENT '创建时间',
				  'update_time' datetime DEFAULT NULL COMMENT '更新时间',
				  PRIMARY KEY ('role_id','template_id') USING BTREE,
      			  FULLTEXT KEY 'title_tags' ('title','tags'),
                  UNIQUE KEY 'userId' ('user_id'),
				  KEY 'createTime' ('create_time') USING BTREE COMMENT '创建时间索引'
				) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色关联表 ';`}
}

//测试所有
func TestCreateParseSql_All(t *testing.T) {
	TestCreateParseSql_Json(t)
	TestCreateParseSql_GoStruct(t)
	TestCreateParseSql_Java(t)
}

//json
func TestCreateParseSql_Json(t *testing.T) {
	jsonRender := NewJsonRenderData()
	DoParse(jsonRender, GetTestCase(), nil)
}

//go
func TestCreateParseSql_GoStruct(t *testing.T) {
	goRender := NewGoStructRenderData()
	DoParse(goRender, GetTestCase(), nil)
}

//java
func TestCreateParseSql_Java(t *testing.T) {
	goRender := NewJavaRenderData()
	DoParse(goRender, GetTestCase(), nil)
}
