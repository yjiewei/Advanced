package main

import (
	"Advanced/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/*
*
关于包的使用：

	1.一个目录下的统计文件归属一个包，package的声明要一致
	2.package声明的包和对应的目录名可以不一致，但习惯上还是写一致的
	3.包可以嵌套，写在多个文件夹里面
	4.同包下的函数不需要导入包，可以直接使用
	5.main包，main()函数所在的包，其他的包不能使用
	6.导入包时，路径要从src下开始写
*/
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
	engine := gin.Default()
	// curl http://localhost:8080/helloworld GET
	engine.GET("/helloworld", func(ctx *gin.Context) {
		// 数组 map list 结构体
		ctx.JSON(200, gin.H{
			"name": "hello world",
		})
	})

	// 路由请求方法
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

	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
