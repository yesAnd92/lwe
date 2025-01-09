package gitcmd

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGitLogSummary(t *testing.T) {
	type args struct {
		detail    bool
		dir       string
		committer string
		start     string
		end       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "",
			args: args{detail: false,
				dir:       "",
				committer: "yesAnd",
				start:     "",
				end:       ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GitLogSummary(tt.args.detail, tt.args.dir, tt.args.committer, tt.args.start, tt.args.end)
		})
	}
}

func Test_logSubmitToAi(t *testing.T) {
	type args struct {
		ctx string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{name: "pageHelper commit log",
			args: args{ctx: `
							Mybatis-PageHelper
							fix 兼容jakarta/javax的ServletRequest
							补充注释信息
							重载一个PageInfo.of方法，支持手动指定查询记录总数返回分页信息
							更新发布日志
							更新release脚本
							6.1.0 更新日志，写文档好麻烦。
							修复错误
							发布6.1.0，jsqlparser直接依赖都是中间接口，可以通过SPI替换默认实现。
							升级jsqlparser版本4.7
							简化pom.xml配置，去掉shade内嵌jsqlparser方式，改为通过外部依赖选择不同的jsqlparser版本，允许自己SPI扩展。
							jsqlparser解析不使用线程池，支持SPI扩展覆盖SqlParser实现
							为了方便使用SqlParser实现，支持SPI方式扩展
							SqlServer分页改为SqlServerSqlParser接口，添加参数 sqlServerSqlParser 覆盖默认值
							OrderByParser提取OrderBySqlParser接口，增加 orderBySqlParser 参数，可以覆盖默认实现。
							OrderByParser静态方法改为普通方法，为后续改接口做准备
							jdk8+后不再需要JSqlParser接口，移除该接口，文档标记该参数
							update README
							发布时打包两个版本，正常版本和standalone版本，standalone版本会内嵌jsqlparser，避免因standalone各个版本不兼容导致的依赖问题。
							内嵌 jsqlparser 以规避版本冲突
							chore: maven-compiler-plugin固定版本以去除警告，并增加构建稳定性`,
			}},
		{name: "lwe commit log",
			args: args{ctx: `
							lwe
							update readme
							add sqllog readme
							add command hit
							add mybatis sql log parse
							update readme
							add sqllog readme
							add command hit
							add mybatis sql log parse
							readme.md add filesserver
							remove pdf command
							update en readme.md
							complete en readme.md
							update no matching commit log prompt
							调整 英文README格式
							add a file server
							add en readme.md
							add env command
							glog add branchs option to determined branch ,limit 50 change to 500
							glog: merge and sort commit log
							get branch commit  log
							add output dir hint of generated file
							add output dir hint of generated file
							Update README.md
							Fsync (#8)
							Gcl command (#7)
							Url command (#5)
							完善Readme.md
							Git command (#3)
							Navi command (#2)
							Initial commit`,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := logSubmitToAi(tt.args.ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(got)
		})
	}
}

func Test_parseResp(t *testing.T) {
	type args struct {
		resp string
	}
	tests := []struct {
		name string
		args args
		want *GitSummaryPromptResp
	}{
		{name: "demo",
			args: args{resp: `
								{
											"repo_summary": [
												{
													"repo": "repository name",
													"summary": [
														"summary1",
														"summary2"
													],
													"summary_cn": [
														"中文总结1",
														"中文总结2"
													]
												}
											]
									}`,
			},
			want: &GitSummaryPromptResp{
				RepoSummary: []RepoSummary{
					{Repo: "repository name",
						Summary:   []string{"summary1", "summary2"},
						SummaryCN: []string{"中文总结1", "中文总结2"},
					},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseResp(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseResp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consoleResult(t *testing.T) {
	type args struct {
		promptResp *GitSummaryPromptResp
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "demo",
			args: args{promptResp: &GitSummaryPromptResp{
				RepoSummary: []RepoSummary{
					{Repo: "repository name",
						Summary:   []string{"summary1", "summary2"},
						SummaryCN: []string{"中文总结1", "中文总结2"},
					},
					{Repo: "repository name2",
						Summary:   []string{"summary1", "summary2"},
						SummaryCN: []string{"中文总结1", "中文总结2"},
					},
				}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			consoleResult(tt.args.promptResp)
		})
	}
}
