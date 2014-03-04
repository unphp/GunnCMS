本程序是基于极好的Go WEB 框架 [Martini](https://github.com/codegangsta/martini)开发的。
***
#### 网站目录基础结构

	|-- mincms
		|-- conf			(配置文件)
		|-- models			(模型)
		|-- public			(公共静态文件)
		|-- templets   		(模板主题)
		|-- utils      		(基础lib)
		|-- readme.md		(自述文件)
		|-- main.go      	(main文件)

### 为什么要制作这个mincms？
GunnCMS是一个针对企业网站设计的内容管理系统。

我想要一个功能简单，并且容易备份和迁移的企业建站程序，所以自己开发了一个，并且也为了练练GOlang，因为我的GOlang一直都是很业余。

它的特点是：

1. 不需要数据库的支持，只需要文件可以运行的环境
1. 只针对小型企业网站网站设计，没有复杂的成员管理和权限设置
1. 使用github的数据api来保持数据文件
1. 使用Markdown撰写文章