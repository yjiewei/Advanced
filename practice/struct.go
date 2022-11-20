package main

import (
	"fmt"
)

/**
结构体：结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
省流：对象
1.定义：
	type struct_variable_type struct {
	   member definition
	   member definition
	   ...
	   member definition
	}
2.声明：
	variable_name := structure_variable_type {value1, value2...valuen}
	variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
3.访问结构体成员：结构体.成员
4.结构体作为函数参数
5.结构体指针 var struct_pointer *Books (var 名字 类型)
	结构体指针参数有点意思 我传给你一个地址，你接收的却是值，跟值传递一样
*/

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	// 创建一个新的结构体 你看像不像对象的默认构造函数新建对象
	fmt.Println(Books{"Go语言", "https://gitee.com/jaysony/go", "学习Go语言", 1})

	// 也可以使用key value的方式，像不像构造函数传入指定参数
	fmt.Println(Books{title: "Go语言", author: "https://gitee.com/jaysony/go", subject: "学习Go语言", bookId: 1})

	// 不传的值会有默认字段
	fmt.Println(Books{title: "Go语言", author: "https://gitee.com/jaysony/go"})

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "https://gitee.com/jaysony/go"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "https://gitee.com/jaysony/go"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookId : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookId : %d\n", Book2.bookId)

	printBook(Book1)

	// 这里的结构体指针参数有点意思 我传给你一个地址，你接收的却是值，跟值传递一样
	/* 打印 Book1 信息 */
	printPointerBook(&Book1)

	/* 打印 Book2 信息 */
	printPointerBook(&Book2)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}

func printPointerBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
