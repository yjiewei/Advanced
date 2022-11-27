package main

import (
	"Advanced/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

/*
*
1.关于包的使用：
	1.一个目录下的统计文件归属一个包，package的声明要一致
	2.package声明的包和对应的目录名可以不一致，但习惯上还是写一致的
	3.包可以嵌套，写在多个文件夹里面
	4.同包下的函数不需要导入包，可以直接使用
	5.main包，main()函数所在的包，其他的包不能使用
	6.导入包时，路径要从src下开始写

2.学习Gin框架，用起来很简洁

*/

// User 字段需要大写
type User struct {
	Id         int               `form:"id" uri:"id"` // tag:form是针对表单的字段映射 还有json
	Name       string            `form:"name" uri:"name"`
	Address    []string          `form:"address"`
	AddressMap map[string]string `form:"addressMap"`
}

func main() {
	// time()
	// file()
	// io()
	// cp()
	// seek()
	// web()
	// templateRendering()
	// setCookie()
	// session()
	// bufferIO()
	// 协程 start 放在main方法中才能正常执行
	//go printNum()
	//for i := 0; i < 100; i++ {
	//	fmt.Printf("这里是主协程打印数字：main - %d\n", i)
	//}
	//fmt.Println("主协程执行结束了。")
	//time2.Sleep(5 * time2.Second)
	// 协程 end
}

func time() {
	utils.Test()
	person := utils.Person{
		Name:  "杨杰炜",
		Email: "yang_7131@163.com",
	}
	fmt.Println("个人姓名：" + person.Name + " 联系方式：" + person.Email)

	utils.TimeFunc()
}

func file() {
	utils.FileFunc()
}

func io() {
	utils.IOFunc()
}

func cp() {
	utils.Copy()
}

func seek() {
	utils.Seek()
}

func web() {
	// 1.获取默认引擎
	engine := gin.Default()
	// curl http://localhost:8080/helloworld GET
	engine.GET("/helloworld", func(ctx *gin.Context) {
		// 数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})

	// 2.路由请求方法
	engine.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(200, "get")
	})

	engine.POST("/post", func(ctx *gin.Context) {
		ctx.JSON(200, "post")
	})

	engine.PUT("/update", func(ctx *gin.Context) {
		ctx.JSON(200, "update")
	})

	engine.DELETE("/delete", func(ctx *gin.Context) {
		ctx.JSON(200, "delete")
	})

	engine.Any("/any", func(ctx *gin.Context) {
		ctx.JSON(200, "any")
	})

	engine.POST("/user/api/:id", func(ctx *gin.Context) {
		ctx.JSON(200, ctx.Param("id"))
	})

	engine.POST("/path/*path", func(ctx *gin.Context) {
		// 获取路径参数
		ctx.JSON(200, ctx.Param("path"))
	})

	// 3.分组路由
	v1 := engine.Group("/v1")
	{
		v1.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v1 find method")
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.GET("find", func(ctx *gin.Context) {
			ctx.JSON(200, "v2 find method")
		})
	}

	// 5.不同参数的后台获取
	// 5.1普通参数
	// http://localhost:8080/user/save?id=1&&name=zhangsan
	engine.GET("/user/save", func(ctx *gin.Context) {
		id := ctx.Query("id")
		name := ctx.Query("name")
		//address, ok := ctx.GetQuery("address")
		address := ctx.DefaultQuery("address", "Beijing")
		ctx.JSON(200, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
			//"ok":      ok,
		})
	})

	engine.GET("/user/saveObj", func(ctx *gin.Context) {
		var user User
		// 下面两个的区别：
		// err := ctx.BindQuery(&user)
		err := ctx.ShouldBindQuery(&user)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(200, user)
	})

	// 5.2数组参数
	// http://localhost:8080/user/array?address=Beijing&&address=Shanghai
	// 如果你想把这个参数丢到user里面，那声明address的时候需要把类型改为[]
	engine.GET("/user/array", func(ctx *gin.Context) {
		address := ctx.QueryArray("address") // 等同于GetQueryArray()
		ctx.JSON(200, address)
	})

	// 5.3map参数
	// http://localhost:8080/user/map?addressMap[home]=Beijing&&addressMap[company]=Shanghai
	engine.GET("/user/map", func(ctx *gin.Context) {
		addressMap := ctx.QueryMap("addressMap") // 等同于GetQueryArray()
		ctx.JSON(200, addressMap)
	})

	// 5.4POST请求获取参数
	engine.POST("/user/save", func(ctx *gin.Context) {
		id := ctx.PostForm("id")
		name := ctx.PostForm("name")
		address := ctx.PostForm("address")
		addressMap := ctx.PostFormMap("addressMap")
		var user User
		_ = ctx.ShouldBind(&user)
		fmt.Println("user:", user)
		ctx.JSON(200, gin.H{
			"id":         id,
			"name":       name,
			"address":    address,
			"addressMap": addressMap,
		})
	})

	// 5.4 json格式的post请求获取参数
	engine.POST("/user/saveJson", func(ctx *gin.Context) {
		var user User
		_ = ctx.ShouldBindJSON(&user)
		fmt.Println("user:", user) // user: {110 杨杰炜 [广东省广州市番禺区] map[company:NETCA name:长腿老头]}
		ctx.JSON(200, user)
	})

	// http://localhost:8080/user/save/10/yangjiewei
	//engine.POST("/user/save/:id/:name", func(ctx *gin.Context) {
	//	id := ctx.Param("id")
	//	name := ctx.Param("name")
	//	ctx.JSON(200, gin.H{
	//		"id":   id,
	//		"name": name,
	//	})
	//})

	// 5.5post请求获取URI上面的参数值
	engine.POST("/user/save/:id/:name", func(ctx *gin.Context) {
		var user User
		_ = ctx.BindUri(&user) //绑定uri上面的参数，user结构体里面要标签上uri
		ctx.JSON(200, user)
	})

	// 5.6post请求获取文件参数
	engine.POST("/user/saveFile", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm() // ctx.FormFile()
		if err != nil {
			log.Fatal(err)
		}
		value := form.Value
		files := form.File
		for _, fileArray := range files {
			for _, v := range fileArray {
				err := ctx.SaveUploadedFile(v, "./"+v.Filename)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		ctx.JSON(200, value)
	})

	// 6.响应不同的数据类型
	engine.POST("/user/response", func(ctx *gin.Context) {
		// 1.字符串类型
		/*ctx.String(http.StatusOK, "this response is string type :%s ", "字符串类型的返回值")
		// 2.返回json
		ctx.JSON(200, gin.H{
			"response": "返回json格式的数据",
		})
		// 3.返回xml
		ctx.XML(200, gin.H{
			"response": "返回xml格式的数据",
		})
		// 4.返回yaml
		ctx.YAML(200, gin.H{
			"response": "返回xml格式的数据",
		})
		// 5.返回文件格式的数据
		ctx.File("./wechat.jpg")*/
		// 6.设置响应的请求头
		ctx.Header("test", "header")
		// 7.设置重定向
		//ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 4.启动项目
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func templateRendering() {
	engine := gin.Default()
	// 1.在加载模板前设置模板函数
	engine.SetFuncMap(template.FuncMap{
		"render": func(str string) template.HTML {
			// 用html渲染函数传入的参数
			return template.HTML(str)
		},
	})
	// 3.加载样式文件 不然如果模板里面使用了会找不到，/css是页面映射的路径，./static/css是实际路径
	engine.Static("/css", "./static/css")
	// 2.加载模板 这里可以加载多个，以及用Glob
	engine.LoadHTMLFiles("./templates/index.tmpl")
	engine.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(200, "index.tmpl", gin.H{
			// "title": "hello template",
			"title": "<a href='http://www.baidu.com'>hello template</a>",
		})
	})
	// 4.
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setCookie() {
	engine := gin.Default()
	// 1.设置cookie
	engine.GET("/cookie", func(ctx *gin.Context) {
		ctx.SetCookie("site_cookie", "this is cookie value", 3600, "/", "localhost", false, true)
		// 2.获取cookie 值
		// ctx.Cookie("site_cookie")
		// 3.删除cookie 把maxAge设置为-1即可
		ctx.SetCookie("site_cookie", "this is cookie value", -1, "/", "localhost", false, true)
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// https://github.com/gin-contrib/sessions
func session() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mySession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8080")
}

// CustomMiddleWare
// 自定义中间件：其实就是一个公有方法，拦截器来的，前置拦截器、后置拦截器，上下文对象设置都是可以的
func CustomMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("这是一个自定义中间件。")
		c.Next()          // 执行下一个中间件
		c.Writer.Status() // 响应状态码
	}
}

func bufferIO() {
	utils.BufferIO()
}

func printNum() {
	for i := 0; i < 100; i++ {
		fmt.Printf("这里是协程打印数字：%d\n", i)
	}
	fmt.Println("协程打印结束。")
}
