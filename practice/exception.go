package main

import "fmt"

/**
错误处理：Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
内置接口：E:/Go/src/builtin/builtin.go:280
*/

// 1.定义一个DivideError结构体
type DivideError struct {
	chuShu    int
	beiChuShu int
}

// 2.实现error接口，这里只是处理除数被除数之间的输出，不是关系运算
// 为什么这个参数类型是指针类型？fixme
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the beiChuShu is zero.
    chuShu: %d
    beiChuJhu: 0
	`
	return fmt.Sprintf(strFormat, de.chuShu)
}

// 3.定义int类型除法运算的函数
// 函数名：Divide 参数两个：chuShu beiChuShu 返回值：result errorMsg
func Divide(chuShu int, beiChuShu int) (result int, errorMsg string) {
	if beiChuShu == 0 {
		data := DivideError{
			chuShu:    chuShu,
			beiChuShu: beiChuShu,
		}
		errorMsg = data.Error()
		return // 返回空
	} else {
		return chuShu / beiChuShu, ""
	}
}

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}

	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
