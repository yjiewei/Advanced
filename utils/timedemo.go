package utils

import (
	"fmt"
	"math/rand"
	"time" // 如果这里报错，说明go语言版本大于goland支持的版本，升级goland可以解决
)

// TimeFunc
/**
时间函数：年月日、时分秒、毫秒微妙纳秒  time.now()
https://pkg.go.dev/time@go1.19.3
*/
func TimeFunc() {
	// 1.获取当前时间 %T可以输出数据类型
	t1 := time.Now()
	fmt.Printf("时间类型是：%T\n", t1)
	fmt.Printf("当前时间是：%s\n", t1)

	t2, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t2)

	t3, err := time.Parse("2006年01月02日 15:04:05", "2022年12月12日 12:12:12")
	if err != nil {
		fmt.Println("转换时间出错", err)
	}
	fmt.Println("转换时间格式：", t3)

	year, month, day := t3.Date()
	fmt.Println("年月日是：", year, month, day)

	hour, minute, second := t3.Clock()
	fmt.Println("时分秒是：", hour, minute, second)

	yearDay := t3.YearDay()
	fmt.Println("每年的第几天：", yearDay)

	t4 := time.Date(2022, time.November, 21, 2, 35, 40, 0, time.UTC)
	fmt.Println("时间戳(秒)：", t4.Unix())      // seconds since 1970
	fmt.Println("时间戳(纳秒)：", t4.UnixNano()) // nanoseconds since 1970

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Println("随机数：", randNum)

}
