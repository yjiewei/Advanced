package main

import (
	"Advanced/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	web()
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

	// 4.启动项目
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
