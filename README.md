lwe是leave work early的缩写，也就是"早点下班"！🤣🤣🤣
它是一个帮助开发者提高工作效率的跨平台命令行工具，或者把它当做go入门学习的项目也是合适的！
总之，欢迎提issue、提好玩或者使用的功能需求，最好能直接PR参与到项目中，大家一起努力，争取早点下班!!! 💪💪💪

## 功能概览

[1.由建表SQL语句转换成Java Bean、Go结构体、Json等](#1)

[2.将SQL语句转换成ElasticSearch查询的DSL语言](#2)

[3.Git增强功能：glog、gl、gcl、gst](#3)

[其它小工具](#4)

## 安装

### 下载编译后的可执行文件

到[release](https://github.com/yesAnd92/lwe/releases)页获取对应平台的版本，在终端上即可使用该二进制文件！

如果你经常使用lwe，更推荐的方式是将二进制文件配置到环境变量中，这样可以随时随地使用二进制文件

更多的安装方式和注意事项，查查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/0.%E5%AE%89%E8%A3%85%E3%80%81%E9%85%8D%E7%BD%AE%E4%BD%BF%E7%94%A8)



## 使用姿势

你可以输入`lwe` 查看lwe命令的使用方式，有哪些子命令及其各自对的功能描述 ，得到使用提示：
```
Usage: lwe [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
  es          Translate SQL to elasticsearch's DSL
  fmt         Generate the specified file based on SQL
  gcl         Update all git repository under the given dir 
  gl          Update all git repository under the given dir 
  glog        Get all git repository commit log under the given dir 
 .....
```
如果你想查看lwe子命令的功能和使用方式，可以使用`-h`参数查看命令的使用帮助
，如：`lwe es -h`


<h3 id="1">1、建表语句生成Java Bean实体、Go 结构体等</h3>

如果我们已经有了表结构，使用建表语句生成对应的实体可以大大减少我们"无脑且重复"工作。
目前支持生成的结构包括Java、Go、Json。

使用方式：

```text
Usage:
lwe fmt [flags]

Examples:
lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/1.%E5%BB%BA%E8%A1%A8SQL%E8%AF%AD%E5%8F%A5%E7%94%9F%E6%88%90%E4%B8%8D%E7%94%A8%E8%AF%AD%E8%A8%80%E6%89%80%E9%9C%80%E5%AE%9E%E4%BD%93)



<h3 id="2">2、SQL语句生成DSL语句</h3>

```bash
lwe es [可选参数] <SQL语句> 
```

这个命令可以帮我们从繁琐的ES查询语法中解脱出来，它可以将sql语句转换成响应的DSL，并且以curl命令的形式输出，这样服务器上也可以方便的使用。
当前版本支持的SQL操作

使用方式：

```text
Usage:
  lwe es [flags]

Examples:
lwe es 'select * from user where age >18' [-p=true]
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/2.%E5%B0%86SQL%E8%AF%AD%E5%8F%A5%E8%BD%AC%E6%8D%A2%E6%88%90ElasticSearch%E6%9F%A5%E8%AF%A2%E7%9A%84DSL%E8%AF%AD%E8%A8%80)


<h3 id="3">3、Git增强功能：glog、gl、gcl、gst</h3>
这里是几个围绕git相关的增强命令，基本都是在原语义上增加了一些跨git仓库的操作



#### glog 增强Git日志功能
查看给定目录下所有git仓库提交日志 
开发人员可能同时维护多个项目或者一个项目中多个模块在不同git仓库，如果有跨仓库查看多个仓库提交日志的需求，glog子命令就派上用场了。

使用方式：

```text
Usage:
  lwe glog [flags]

Examples:
lwe glog [git repo dir] [-a=yesAnd] [-n=50] [-s=2023-08-04] [-e=2023-08-04]
```


#### gl 增强拉取代码功能
拉取给定目录下的所有git仓库最新代码(使用的git pull --rebase的方式)

Git增强功能详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/3.Git%E5%A2%9E%E5%BC%BA%E5%8A%9F%E8%83%BD)


使用方式：
```text
Usage:
  lwe gl [flags]

Examples:
lwe gl [git repo dir]
```

#### gcl 增强git clone功能
使用方式：
```text
Usage:
  lwe gcl [flags]

Examples:
lwe gcl gitGroupUrl [dir for this git group] -t=yourToken
```

#### gst 查看指定目录下所有git仓库状态
查看给定目录下的所有git仓库状态

使用方式：
```text
Usage:
  lwe gst [flags]

Examples:
lwe gst [your git repo dir]
```


<h3 id="4">其它小工具</h3>
一些非常实用的功能

<h4>格式化请求url</h4>
有时请求的url很长，不利于我们找到目标参数，可以使用url命令进行格式化，增加请求的可读性

使用方式：

```text
Usage:
  lwe url [flags]

Examples:
lwe url yourUrl
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E6%A0%BC%E5%BC%8F%E5%8C%96%E8%AF%B7%E6%B1%82url)


<h4>获取Navicat连接配置中的密码</h4>
如果想从Navicat保存的连接中获取对应数据库的用户名/密码，可以使用ncx文件，ncx文件是Navicat导出的连接配置文件，但ncx中的密码是一个加密后的十六进制串，使用ncx命令可以获取对应的明文

使用方式：

```text
Usage:
lwe ncx [flags]

Examples:
lwe ncx ncx-file-path
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E8%8E%B7%E5%8F%96navicat%E8%BF%9E%E6%8E%A5%E9%85%8D%E7%BD%AE%E4%B8%AD%E7%9A%84%E5%AF%86%E7%A0%81)

<h4>同步两个目录下文件</h4>
如果你有备份文件的习惯，这个工具可能会帮到你，它可以将源目录文件下的新增的文件同步到备份目录，省去了你逐层文件夹逐个文件去手动同步。

使用方式：
```text
Usage:
lwe fsync [flags]

Examples:
lwe fsync sourceDir targetDir [-d=true]
```

详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E5%90%8C%E6%AD%A5%E4%B8%A4%E4%B8%AA%E7%9B%AE%E5%BD%95%E4%B8%8B%E6%96%87%E4%BB%B6)



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
