## 1.菜鸟教程
- 花了几天时间把菜鸟教程上的Go基础语法过了一遍，在helloworld/practice目录下，教程地址 
  https://www.runoob.com/go/go-tutorial.html
- 接下来会看看web相关的教程吧，冲就完事了

## 2.Go进阶
- 原先基础部分的代码示例复制到practice目录下了，这里的包及文件不是很规范，但方便
- https://www.bilibili.com/video/BV1MA411K7Yh

## 3.不知记哪的笔记 
- init()、main()函数
  - 相同点：
    1. 两个函数在定义时不能有任何的参数和返回值；
    2. 该函数只能由go程序自动调用，不可以被引用。
  - 不同点
    1. init函数可以应用于任意包中，且可以重复定义多个；
    2. main函数只能用于main包中，且只能定义一个。
  - 两个函数的执行顺序
    1. 在main包中的go文件默认总是会被执行；
    2. 对同一个go文件的init函数调用顺序是从上到下的；
    3. 对同一个package中的不同文件，将文件名按字符串进行从小到大排序，之后顺序调用各文件助攻的init函数；
    4. 对于不同的package，如果不相互依赖的话，按照main包中的import顺序调用其包中的init函数。
    5. 如果package存在依赖，调用顺序为最后被依赖的最先被初始化 (常量变量会在init函数之前初始化)
    6. 避免循环依赖，一个包可以被多个包引入，但不会多次初始化
- 为了执行包的初始化函数而导包，导入的时候可以对包起别名 _
- go允许import不同代码库的代码，对于import要导入的外部的包，可以使用go get命令去下来放到GOPATH下，但是用模块管理之后要怎么做？

## 4.Gin
- 视频地址：https://www.bilibili.com/video/BV1wG4y1Z7Wo
- 官网地址：https://gin-gonic.com/zh-cn/docs/
- go work init 
- go mod init [module path]
- go get -u github.com/gin-gonic/gin
- 类似Python的导包，导入之后就能使用了，不一定只能用这个框架
- 书写URI的时候，我们不需要关心schema和authority这两部分，主要通过path和query两部分书写来进行资源的定位，比如静态URI，路径参数
- 处理函数，type HandlerFunc func(*Context) ，通过上下文参数，获取HTTP请求参数，响应HTTP请求，比如 `func(ctx *gin.Context) { ctx.JSON(200, ctx.Param("path"))}`
- 分组路由，比如分成版本、服务之间的区分 `v1 := engine.Group("/v1")`  `user := engine.Group("/user")`

## X.碎碎念
1. Advanced意思是进阶，想要取得进步继续努力；
2. 学Go最初的想法就是担心Java太卷，想换个赛道；
3. 20221123广州疫情形势依旧很严峻，天天闷家里快发霉了；