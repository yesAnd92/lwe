## lwe
lwe是leave work early的缩写，也就是"早点下班"！🤣🤣🤣
它是一个帮助开发者提高工作效率的跨平台命令行工具，当然你把它当做go入门学习的项目也是合适的！
总之，欢迎提issue、提好玩或者使用的功能需求，最好能直接PR参与到项目中，大家一起努力，争取早点下班!!! 💪💪💪


## 功能

[Git增强多仓库操作能力：glog、gl、gcl、gst](#git)

[由建表SQL语句转换成Java Bean、Go结构体、Json等](#fmt)

[将SQL语句转换成ElasticSearch查询的DSL语言](#es)

[PDF工具：合并多个图片或者PDF、截取PDF指定页](#pdf)

[其它小工具](#other)
- 获取Navicat连接配置中的密码
- 同步两个目录下文件
- 格式化请求url


## 安装
### 下载编译后的可执行文件

到[release](https://github.com/yesAnd92/lwe/releases)页获取对应平台的版本，在终端上即可使用该二进制文件！

当然，更推荐的方式是将二进制文件配置到环境变量中，这样可以随时随地使用二进制文件

更多的安装方式和注意事项，查查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/0.%E5%AE%89%E8%A3%85%E3%80%81%E9%85%8D%E7%BD%AE%E4%BD%BF%E7%94%A8)



## 使用

你可以输入`lwe` 查看lwe命令的使用方式，有哪些子命令及其各自对的功能描述。

如果对某个子命令感兴趣，可以使用`-h`参数查看命令的使用示例 ，如：`lwe glog -h`


<h3 id="git">Git增强多仓库操作能力：glog、gl、gcl、gst</h3>
这里是几个围绕git相关的增强命令，基本都是在原语义上增加了一些跨git仓库的操作

Git增强功能详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/3.Git%E5%A2%9E%E5%BC%BA%E5%8A%9F%E8%83%BD)

#### glog 增强Git日志功能
查看给定目录下所有git仓库提交日志
开发人员通常会在多个git仓库下工作，经常会有同时查看多个git仓库提交日志的需求，glog子命令就派上用场了。

使用方式：

```text
lwe glog [git repo dir] [-a=yesAnd] [-n=50] [-s=2023-08-04] [-e=2023-08-04]
```


#### gl 增强拉取代码功能
拉取给定目录下的所有git仓库最新代码(使用的git pull --rebase的方式)



使用方式：
```text
lwe gl [git repo dir]
```

#### gcl 增强git clone功能
使用方式：
```text
lwe gcl gitGroupUrl [dir for this git group] -t=yourToken
```

#### gst 查看指定目录下所有git仓库状态
查看给定目录下的所有git仓库状态

使用方式：
```text
lwe gst [your git repo dir]
```
***

<h3 id="fmt">建表语句生成Java Bean实体、Go 结构体等</h3>

如果我们已经有了表结构，使用建表语句生成对应的实体可以大大减少我们"无脑且重复"工作。
目前支持生成的结构包括Java、Go、Json。

使用方式：

```text
lwe fmt sql-file-path [-t=java|go|json] [-a=yesAnd]
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/1.%E5%BB%BA%E8%A1%A8SQL%E8%AF%AD%E5%8F%A5%E7%94%9F%E6%88%90%E4%B8%8D%E7%94%A8%E8%AF%AD%E8%A8%80%E6%89%80%E9%9C%80%E5%AE%9E%E4%BD%93)

***

<h3 id="es">SQL语句生成DSL语句</h3>

```bash
lwe es [可选参数] <SQL语句> 
```

这个命令可以帮我们从繁琐的ES查询语法中解脱出来，它可以将sql语句转换成响应的DSL，并且以curl命令的形式输出，这样服务器上也可以方便的使用。
当前版本支持的SQL操作

使用方式：

```text
lwe es 'select * from user where age >18' [-p=true]
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/2.%E5%B0%86SQL%E8%AF%AD%E5%8F%A5%E8%BD%AC%E6%8D%A2%E6%88%90ElasticSearch%E6%9F%A5%E8%AF%A2%E7%9A%84DSL%E8%AF%AD%E8%A8%80)

***

<h3 id="pdf">PDF工具：合并多个图片或者PDF、截取PDF指定页</h3>

对PDF进行简单的编辑还是挺常用的功能，比如把几个PDF、图片合并成一个PDF，从一个PDF中抽离出指定页等，但是在很多办公软件上收费的功能，lwe提供了简单编辑PDF的能力。

#### pdfm 合并PDF或者图片
按照指定的文件顺序，将多个PDF、图片文件合并成一个PDF文件

使用方式：
```text
lwe pdfm out.pdf in1.pdf,in2.jpg,*.png,in3.pdf ...
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/PDF%E5%B7%A5%E5%85%B7%EF%BC%9A%E5%90%88%E5%B9%B6%E5%A4%9A%E4%B8%AA%E5%9B%BE%E7%89%87%E6%88%96%E8%80%85PDF%E3%80%81%E6%88%AA%E5%8F%96PDF%E6%8C%87%E5%AE%9A%E9%A1%B5#pdfm-%E5%90%88%E5%B9%B6pdf%E6%88%96%E8%80%85%E5%9B%BE%E7%89%87)


#### pdfc 从PDF中截取指定的页
按照指定的页码，从PDF中截取对应的页生成PDF文件

使用方式：
```text
lwe pdfc [-m] in.pdf outDir 2,3,5,7-9,15 ...
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/PDF%E5%B7%A5%E5%85%B7%EF%BC%9A%E5%90%88%E5%B9%B6%E5%A4%9A%E4%B8%AA%E5%9B%BE%E7%89%87%E6%88%96%E8%80%85PDF%E3%80%81%E6%88%AA%E5%8F%96PDF%E6%8C%87%E5%AE%9A%E9%A1%B5#pdfc-%E4%BB%8Epdf%E4%B8%AD%E6%88%AA%E5%8F%96%E6%8C%87%E5%AE%9A%E7%9A%84%E9%A1%B5)

***

<h3 id="other">其它小工具</h3>
一些非常实用的功能

<h4>格式化请求url</h4>
有时请求的url很长，不利于我们找到目标参数，可以使用url命令进行格式化，增加请求的可读性

使用方式：

```text
lwe url yourUrl
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E6%A0%BC%E5%BC%8F%E5%8C%96%E8%AF%B7%E6%B1%82url)


<h4>获取Navicat连接配置中的密码</h4>
如果想从Navicat保存的连接中获取对应数据库的用户名/密码，可以使用ncx文件，ncx文件是Navicat导出的连接配置文件，但ncx中的密码是一个加密后的十六进制串，使用ncx命令可以获取对应的明文

使用方式：

```text
lwe ncx ncx-file-path
```
详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E8%8E%B7%E5%8F%96navicat%E8%BF%9E%E6%8E%A5%E9%85%8D%E7%BD%AE%E4%B8%AD%E7%9A%84%E5%AF%86%E7%A0%81)

<h4>同步两个目录下文件</h4>
如果你有备份文件的习惯，这个工具可能会帮到你，它可以将源目录文件下的新增的文件同步到备份目录，省去了你逐层文件夹逐个文件去手动同步。

使用方式：
```text
lwe fsync sourceDir targetDir [-d=true]
```

详细使用说明，可以查阅[Wiki](https://github.com/yesAnd92/lwe/wiki/%E5%85%B6%E5%AE%83%E5%B0%8F%E5%B7%A5%E5%85%B7#%E5%90%8C%E6%AD%A5%E4%B8%A4%E4%B8%AA%E7%9B%AE%E5%BD%95%E4%B8%8B%E6%96%87%E4%BB%B6)



## 声明
1.使用[spf13/cobra](github.com/spf13/cobra)库来方便的构建命令行工具

2.es子命令实现借助了[sqlparser](github.com/xwb1989/sqlparser)库来解析SQL语句，一个库很优秀的解析SQL库

3.sql转换成dsl，曹大的[elasticsql](https://github.com/cch123/elasticsql)项目已经是一个很成熟好用的轮子了，lwe也大量借鉴了它的实现思路；没直接调用这个库的原因是想自己练手，同时后续增减功能也更加灵活

4.git增强命令结果输出时使用了[go-pretty](https://github.com/jedib0t/go-pretty)库来表格化提交信息

5.pdf命令是在[pdfcpu](https://github.com/pdfcpu/pdfcpu)基础上进行的封装

## RoadMap
- fmt 根据需求支持更多类型的转换
- es 按需增加对insert、update、delete
  ......

## 开源协议

[MIT License](https://github.com/yesAnd92/lwe/blob/main/LICENSE)
