lwe是leave work early的缩写，也就是"早点下班"！🤣🤣🤣
它是一个帮助开发者提高工作效率的跨平台命令行工具，或者把它当做go入门学习的项目也是合适的！
总之，欢迎提issue、提好玩或者使用的功能需求，最好能直接PR参与到项目中，大家一起努力，争取早点下班!!! 💪💪💪

## 功能概览

[1.支持由建表SQL语句转换成Java Bean、Go结构体、Json等](#1)

[2.支持将查询SQL语句转换成ElasticSearch查询的DSL语言](#2)

[3.获取给定值的md5值](#3)

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
help        Help about any command
md5         Get a md5 for the given value or  a random md5 value
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

> 一个SQL文件中支持多个创建语句;

> 另外，实际使用中最好使用比如Navicat等工具导出的建表语句，识别率会更高。自己手写的可能由于语法或者拼写错误导致错误识别！

你可以使用以下命令来生成Java的实体Bean

```bash
lwe fmt -t=java -a=yesAnd user.sql
```

其中：

`-a, --author string   Comment for author information will be added to the generated file`,该参数用于指定生成文件的注释中作者的信息。\
`-t, --target string   The type[java|json|go] of generate the sql (default "java")`,该参数用于指定生成文件类型，目前支持\[java|go|json],默认值是java，即生成Java Bean。

执行命令后会在`lwe-generate-file`目录下生成相应的文件`StudentInfo.java`,内容如下：

```java
import java.util.Date;
import java.util.List;
import java.io.Serializable;
import javax.persistence.Column;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.GeneratedValue;

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

    public Integer getId() {
        return id;
    }
    public void setId(Integer id) {
        this.id = id;
    }
    //仅展示格式，省略部分字段
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

如果经常使用ElasticSearch的dsl在本地做一些查询并且又不想记忆繁琐的语法，这个命令就派上用场了，毕竟大家对SQL语句还是信手拈来的。

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

```json
// POST /user/_search
{
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
}
```
<h3 id="3">3、获取给定值的md5值</h3>
这个命令非常的简单，返回给定值的md5值，如果未给定值则随机返回一个md5值

```bash
lwe md5 [给定的串]
```
如：
```bash
lwe md5 yesAnd
```


## 说明
1.使用[spf13/cobra](github.com/spf13/cobra)库来方便的构建命令行工具

2.es子命令实现借助了[sqlparser](github.com/xwb1989/sqlparser)库来解析SQL语句，一个库很优秀的解析SQL库

3.sql转换成dsl，曹大的[elasticsql](https://github.com/cch123/elasticsql)项目已经是一个很成熟好用的轮子了，lwe也大量借鉴了它的实现思路；没直接调用这个库的原因是想自己练手，同时后续增减功能也更加灵活

## RoadMap
- fmt 根据需求支持更多类型的转换
- es 按需增加对insert、update、delete
......

## 开源协议

[MIT License](https://github.com/yesAnd92/lwe/blob/main/LICENSE)
