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
- 关于ioutil包 `Deprecated: As of Go 1.16, the same functionality is now provided by package io or package os, and those implementations should be preferred in new code. See the specific function documentation for details.`
- 并发性Concurrency
  - 多任务：我的电脑一边跑goland，一边放哔哩哔哩视频，一边挂着微信...因为CPU的速度非常快。
  - 什么是并发：比如我们浏览器在下载文件和请求网页，单核处理器中，上下文来回切换，一会下载一会加载网页数据，这就是并发，从不同时间点开始，执行周期重叠。**如果是多核，他们同时在不同内核中运行，那他们就是并行。**
  - 什么是并行：我一边听歌，一边写代码，一边抖腿...，并行不一定比并发快，因为并行需要组件间的相互通信。
  - 进程、线程、协程(轻量级线程)：我开了微信(进程)，微信里面有朋友圈新消息提醒还能聊天(线程)
  - 协程：Coroutine，由用户控制，通常跟子函数放在一起比较，协程和多线程相比，其优势体现在：协程执行效率极高。因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程相比，线程数量越多，协程的性能优势越明显。
- Go语言对于并发的实现就是靠协程，Goroutine
- Go runtime包 `Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines. It also includes the low-level type information used by the reflect package; see reflect's documentation for the programmable interface to the run-time type system.`

## 4.Goroutine
- Goroutines在线程上的优势：
  - 资源消耗低，只是堆栈大小的几个kb，可以根据程序需要增长收缩，而线程的堆栈是指定并且固定的
  - Goroutines被多路复用到较少的OS线程。fixme:不懂什么意思
  - 当使用Goroutines访问共享内存时，通过设计的通道可以防止竞态条件发生，通过可以认为是Goroutines通信的管道。
- 封装main函数的Goroutine被称为主Goroutine，职责以及工作流程
  - 设定每一个Goroutine所能申请的栈空间最大尺寸，栈空间超过这个则会stack overflow
  - 创建defer语句，主Goroutine退出时的后续处理
  - 启动专用于在后台清扫内存垃圾的Goroutine，并设置GC可用的标识
  - 执行main包中的init函数
  - 执行main函数
- Goroutine不建议写返回值
- main中的Goroutine终止了，程序将被终止
- Go语言的并发模型可以重点多看，文字描述内容比较多，https://www.bilibili.com/video/BV1MA411K7Yh/?p=16
  - 线程模型：1.内核级线程模型 2.用户级线程模型 3.，两级线程模型(Goroutine)
- 临界资源安全问题
  - 临界资源：指并发环境中多个进程、线程、协程共享的资源，如果处理不当，往往会导致数据不一致问题。
  - `go run -race main.go` 这个命令可以显示程序执行路径
  - 可以用 sync.WaitGroup、互斥锁 去解决
  - https://pkg.go.dev/sync#Mutex
  - https://pkg.go.dev/sync#WaitGroup
  - 用互斥锁对资源很浪费，可以选择读写锁 https://pkg.go.dev/sync#RWMutex
  - 通道channel，goroutines之间通信的管道，数据可以从一端流向另一端通过管道去接收。
  - “不要通过共享内存来通信，而应该通过通信来共享内存”

## 5.chan
- 通道
  - goroutines之间通信的管道，通道的发送和接收必须处在不同的goroutine中
  - 声明通道和定义一个变量一样 声明：`var 通道名 chan 数据类型` 创建：`通道名 = make(chan 数据类型)`
  - 可以使用通道传递数据，也能起到阻塞的作用，本身channel就是同步的，同一时间只能有一条goroutine来操作。
  - 阻塞：不管你是读还是写都可能被阻塞，读的时候没有数据就会被阻塞，写的时候已有数据没有被读取也会阻塞。默认读写都是阻塞的，前提是没有缓存
  - 死锁：只有读操作或者只有写操作就会出现死锁
  - 如何关闭通道？close(chan) 会通知接收方不会有更多数据被发送到channel上了 `v,ok := <-channel` 
  - 通道上如何使用范围循环？通过range去判断接收的数据，不需要通过ok值
  - 缓冲通道、定向通道
  - time包中的通道相关函数 定时器 time.NewTimer(3*time.Second)
- select语句：类似于switch语句
  - 每个case都必须是一个通信
  - 所有channel表达式都会被求值
  - 所有被发送的表达式都会被求值
  - 如果有多个可以执行的case，会随机执行一个
  - 如果没有case符合条件，执行default或阻塞
  - 可以用来监听通道的数据流动
- Go语言的CSP模型：通信顺序进程，是一种并发编程模型

## 6.反射reflect
- 概念：反射是指计算机程序在运行时可以访问、检测和修改它本身状态或行为的能力。
- 反射的使用场景:
  - 有时你需要编写一个函数，但是不知道传给你的参数类型是什么，可能是没约定好；也可能是传入的类型很多，这些类型并不能统一表示。这是反射可以用上；
  - 有时候需要根据某些条件决定调用哪个函数，比如根据用户的输入去决定，这时需要对函数和函数的参数进行反射，在运行期间动态地执行函数。
- 反对反射的理由：
  - 代码可读性不够；
  - 关于反射的代码可能需要运行很久，才会出错，并且直接panic可能会造成严重后果；
  - 对性能影响比较大，比正常代码运行速度慢一到两个数量级。
- https://pkg.go.dev/reflect
- reflect.TypeOf()是获取pair中的type，reflect.ValueOf()获取pair中的value

## 7.Gin
- 视频地址：https://www.bilibili.com/video/BV1wG4y1Z7Wo
- 官网地址：https://gin-gonic.com/zh-cn/docs/
- go work init 
- go mod init [module path]
- go get -u github.com/gin-gonic/gin
- 类似Python的导包，导入之后就能使用了，不一定只能用这个框架
- 书写URI的时候，我们不需要关心schema和authority这两部分，主要通过path和query两部分书写来进行资源的定位，比如静态URI，路径参数
- 处理函数，type HandlerFunc func(*Context) ，通过上下文参数，获取HTTP请求参数，响应HTTP请求，比如 `func(ctx *gin.Context) { ctx.JSON(200, ctx.Param("path"))}`
- 分组路由，比如分成版本、服务之间的区分 `v1 := engine.Group("/v1")`  `user := engine.Group("/user")`
- 获取Get请求和Post请求带的参数值，包括不同的数据类型
- 响应的不同数据类型
- 模板渲染，看样子是集成了前端
- cookie & session，session我们用中间件去处理`go get github.com/gin-contrib/sessions`
- 基于Redis作为存储引擎保存session `go get github.com/gin-contrib/sessions/redis` 官网上都有地址
- 中间件（也可以自定义中间件），在请求和响应之间的特殊函数，每个中间件执行不同的功能，一个中间件执行完就下一个。
  - 使用场景：请求限速、api接口签名处理、权限校验、统一错误处理
  - 在Java中其实对应的是拦截器
  ```go
    // Default returns an Engine instance with the Logger and Recovery middleware already attached.
    func Default() *Engine {
      debugPrintWARNINGDefault()
      engine := New()
      // 这里用了两个中间件 日志和Recovery
      engine.Use(Logger(), Recovery())
      return engine
    }
  ```

## X.碎碎念
1. Advanced意思是进阶，想要取得进步继续努力；
2. 学Go最初的想法就是担心Java太卷，想换个赛道；
3. 20221123广州疫情形势依旧很严峻，天天闷家里快发霉了；
4. 20221128女朋友闹分手，疫情原因也见不上，这几天也没怎么写代码；