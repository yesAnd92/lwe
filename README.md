lwe是leave work early的缩写，也就是"早点下班"！🤣🤣🤣
它是一个帮助开发者提高工作效率的跨平台命令行工具，或者把它当做go入门学习的项目也是合适的！
总之，欢迎提issue、提好玩或者使用的功能需求，最好能直接PR参与到项目中，大家一起努力，争取早点下班!!! 💪💪💪

## 功能概览

[1.由建表SQL语句转换成Java Bean、Go结构体、Json等](#1)

[2.将SQL语句转换成ElasticSearch查询的DSL语言](#2)

[3.获取Navicat连接配置中的密码](#3)

[4.增强Git日志功能：查看给定目录下所有git仓库提交日志](#4)

[5.其它小工具](#5)

## 安装

### 下载编译后的可执行文件

到[release](https://github.com/yesAnd92/lwe/releases)页获取对应平台的版本，在终端上即可使用该二进制文件！

当然，也可以将二进制文件配置到环境变量中，这样可以随时随地使用二进制文件

#### Mac平台添加到环境变量的方法:

```bash
cp <下载的lwe文件路径> /usr/local/bin
```
将下载好的lwe_Mac文件直接cp到`/usr/local/bin`目录下，即可在终端下使用了
> 最好将lwe_Mac重名成lwe，方便使用

>执行cp命令时，通常需要提升权限，可以在`cp`命令前增加`sudo`来解决


#### Win平台添加到环境变量的方法:
可以参照maven的配置方法

### Homebrew

Mac平台后续支持从Homebrew直接安装

### 使用源码编译

源码下载到本地，使用go build或者make命令进行编译，生成可用的二进制文件，make 的其它使用命令可以参见makefile文件

> 本地编译需要go环境



## 使用姿势

```
Usage: lwe [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
es          Translate SQL to elasticsearch's DSL
fmt         Generate the specified file based on SQL
glog        Get all git repository commit log under the given dir 
help        Help about any command
md5         Get a md5 for the given value or  a random md5 value
ncx         Decrypt password of connection in .ncx file
version     Print the version number of lwe
```
### help

lwe命令以及其子命令都可以使用`-h`参数查看命令的使用帮助


<h3 id="1">1、建表语句生成Java Bean实体、Go 结构体等</h3>

```bash
lwe fmt [可选参数] <建表语句的文件路径> 
```

例如当前目录下，有个user.sql建表语句，内容如下：

```sql
CREATE TABLE 'student_info' (
				  'id' int(11) NOT NULL AUTO_INCREMENT COMMENT '用户编号,学号',
				  'class_id' varchar(255) NOT NULL COMMENT '班级id',
				  'user_name' varchar(255) NOT NULL COMMENT '用户名',
				  'status' tinyint(1) NOT NULL COMMENT '状态',
				  'create_time' datetime NOT NULL COMMENT '创建时间',
				  PRIMARY KEY ('id') 
				) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学生信息';
```

> 一个SQL文件中支持多个创建语句批量生成目标文件;

> 另外，实际使用中最好使用比如Navicat等工具导出的建表语句，识别率会更高。自己手写的可能由于语法或者拼写错误导致错误识别！

你可以使用以下命令来生成Java的实体Bean

```bash
lwe fmt -t=java -a=yesAnd user.sql
```

其中：

`-a, --author string   Comment for author information will be added to the generated file`,可选参数，该参数用于指定生成文件的注释中作者的信息。\
`-t, --target string   The type[java|json|go] of generate the sql (default "java")`,该参数用于指定生成文件类型，目前支持[java|go|json],默认值是java，即生成Java Bean。

执行命令后会在`lwe-generate-file`目录下生成相应的文件`StudentInfo.java`,内容如下：

```java
//省略部分字段仅做展示
import java.util.Date;
import javax.persistence.Id;
...

/**
 * @Description 学生信息
 * @Author  yesAnd
 * @Date  2023/04/17 10:28
 */
@Table ( name ="student_info" )
public class StudentInfo implements Serializable {

    private static final long serialVersionUID = 1L;
    
    @Id 
    @Column(name = "id" )
    private Integer id;	//用户编号，学号

    public Integer getId() { return id;}
    
    public void setId(Integer id) {  this.id = id;}
```

同样的，指定`-t=go`生成对应的结构体：

```go
//StudentInfo 学生信息
type StudentInfo struct {
	Id         int       `gorm:"primary_key;" json:"id"` //用户编号，学号
	ClassId    string    `gorm:"" json:"classId"`        //班级id
	UserName   string    `gorm:"" json:"userName"`       //用户名
	Status     int64     `gorm:"" json:"status"`         //状态
	CreateTime time.Time `gorm:"" json:"createTime"`     //创建时间

}
```

###

<h3 id="2">2、SQL语句生成DSL语句</h3>

```bash
lwe es [可选参数] <SQL语句> 
```

这个命令可以帮我们从繁琐的ES查询语法中解脱出来，它可以将sql语句转换成响应的DSL，并且以curl命令的形式输出，这样服务器上也可以方便的使用。
当前版本支持的SQL操作

✅select

*   [x] and&#x20;
*   [x] or&#x20;
*   [x] \=&#x20;
*   [x] <  <=  >  >=
*   [x] in  not in
*   [x] like   not like
*   [x] order by&#x20;
*   [x] limit
*   [ ] group by
*   [ ] join&#x20;
*   [ ] having support

❌ update

❌ delete

❌ insert

es子命令的使用也非常简单，例如：

```bash
lwe es -p 'select * from user where age >18  order by create_time desc  limit 10,10'
```

其中：

`-p, --pretty   Beautify DSL`,该参数用于美化生成的DSL结果

生成的结果如下：

```bssh
curl -XPOST -H "Content-Type: application/json" -u {username}:{password} {ip:port}/user/_search?pretty -d '{
  "from": 10,
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "age": {
              "gt": "18"
            }
          }
        }
      ]
    }
  },
  "size": 10,
  "sort": [
    {
      "create_time": "desc"
    }
  ]
}'
```

<h3 id="3">3、获取Navicat连接配置中的密码</h3>
如果想从Navicat保存的连接中获取对应数据库的用户名/密码，可以使用ncx文件，ncx文件是Navicat导出的连接配置文件，但ncx中的密码是一个加密后的十六进制串，使用ncx命令可以获取对应的明文:

```bash
lwe ncx <ncx文件路径>
```
> Navicat导出连接的步骤：file->export connections->勾选 export password选项->确定

如： 导出一个名为local-mysql的连接demo.ncx，内容是：
```xml
<!--仅节选几个重要字段作为说明展示-->
<Connections Ver="1.5">
  <Connection ConnectionName="local-mysql"  ConnType="MYSQL"  Host="127.0.0.1" Port="3306" UserName="root" Password="B75D320B6211468D63EB3B67C9E85933" />
</Connections>
```
使用ncx命令：
```bash
lwe ncx ./demo.ncx
```
输出结果：
```text
-----------local-mysql-----------
DB type:  MYSQL
Connection host: 127.0.0.1
Connection port: 3306
Connection username: root
Connection password: This is a test
```

<h3 id="4">4、增强Git日志功能：查看给定目录下所有git仓库提交日志</h3>
开发人员可能同时维护多个项目或者一个项目中多个模块在不同git仓库，如果有跨仓库查看多个仓库提交日志的需求，glog子命令就派上用场了。

```bash
lwe glog  <仓库所在目录>  [可选参数] 
```

> 查询结果对应的是每个git仓库当前使用分支的提交记录

> 如果未指定目录，则在当前目录下搜寻git仓库,另，如果目录层级过深，可能会影响性能

如：写周报时，需要查看自己近一周在哪些仓库提交了哪些代码,来辅助我写总结，假定我的工作目录在/Users/yesand/work/

```bash
lwe glog /Users/yesand/work/  -a=yesand -f=false -n=20 -s=2023-05-15 -e=2023-05-19
```
其中：

`-a, --author string`,可选参数，该参数用于指定提交者，未指定查询所有提交者。\
`-f, --file bool`,可选参数，该参数决定将查询结果写到文件中，默认在控制台输出。\
`-n, --recentN int16`,可选参数，该参数指定每个仓库查询最近N条的提交记录。\
`-s, --start string`,可选参数，该参数指定筛选提交记录的开始日期，格式：'yyyy-MM-dd'。\
`-e, --end string`,可选参数，该参数指定筛选提交记录的结束日期，格式：'yyyy-MM-dd'。\

结果:示例

```text
#1 Git Repo >> /Users/yesand/work/lwe
+---------+--------+-----------------------------------------+---------------------+
| HASH    | AUTHOR | COMMIT                                  | TIME                |
+---------+--------+-----------------------------------------+---------------------+
| bf67fcd | yesand | 完善命令提示&交互提示                      | 2023-05-19 17:21:34 |
| 3739c60 | yesand | 优化build后的二进制文件大小                 | 2023-05-19 09:44:14|
| 7a2ca47 | yesand | 以表格形式输出提交记录更换为go-pretty库      | 2023-05-19 09:21:26 |
+---------+--------+-----------------------------------------+---------------------+

#2 Git Repo >> /Users/yesand/work/xxx
...
```

<h3 id="5">5、其它小工具</h3>
一些小的功能

<h4>格式化请求url</h4>
有时请求的url很长，不利于我们找到目标参数，可以使用url命令进行格式化，增加请求的可读性
示例：
```bash
lwe url  http://api.demo.com/api/user/getList?platform=ios&signature=bd7e1fd4e65e8199fd817006e709fb33&currentTimeMillis=1685673384000&pageNum=1
```
格式化结果：
```text
Host: api.demo.com
Path: /api/user/getList
-----------------------
pageNum                 1
platform                ios
signature               bd7e1fd4e65e8199fd817006e709fb33
currentTimeMillis       1685673384000
```

> 某些bash下请求url需要用' '引起来才能正常使用


## 说明
1.使用[spf13/cobra](github.com/spf13/cobra)库来方便的构建命令行工具

2.es子命令实现借助了[sqlparser](github.com/xwb1989/sqlparser)库来解析SQL语句，一个库很优秀的解析SQL库

3.sql转换成dsl，曹大的[elasticsql](https://github.com/cch123/elasticsql)项目已经是一个很成熟好用的轮子了，lwe也大量借鉴了它的实现思路；没直接调用这个库的原因是想自己练手，同时后续增减功能也更加灵活

4.glog结果输出时使用了[go-pretty](https://github.com/jedib0t/go-pretty)库来表格化提交信息
## RoadMap
- fmt 根据需求支持更多类型的转换
- es 按需增加对insert、update、delete
  ......

## 开源协议

[MIT License](https://github.com/yesAnd92/lwe/blob/main/LICENSE)
