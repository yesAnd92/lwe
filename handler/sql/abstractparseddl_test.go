package sql

import (
	"os"
	"path"
	"testing"
)

var (
	sqlTextArr = []string{
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
)

func TestDoParse(t *testing.T) {
	type args struct {
		parse      IParseDDL
		sqlTextArr []string
		args       map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "json",
			args: args{
				parse:      NewJsonRenderData(),
				sqlTextArr: sqlTextArr,
				args:       nil,
			},
			want: []string{},
		},
		{
			name: "java",
			args: args{
				parse:      NewJavaRenderData(),
				sqlTextArr: sqlTextArr,
				args:       nil,
			},
			want: []string{"ClassInfo.java", "StudentInfo.java", "TemplateRole.java"},
		},
		{
			name: "go",
			args: args{
				parse:      NewGoStructRenderData(),
				sqlTextArr: sqlTextArr,
				args:       nil,
			},
			want: []string{"lwe_struct.go"},
		},
	}
	defer func() {
		os.RemoveAll(GENERATE_DIR)
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoParse(tt.args.parse, tt.args.sqlTextArr, tt.args.args)
			for _, p := range tt.want {
				if f, err := os.Stat(path.Join(GENERATE_DIR, p)); err != nil || f.Size() == 0 {
					t.Errorf("file >>> %s is not except", p)
				}
			}
		})
	}
}
