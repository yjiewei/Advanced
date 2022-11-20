package main

import "fmt"

/*
*

	切片
	1.概念：Go 语言切片是对数组的抽象。
	       Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，
		   与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
	2.声明切片：
		var identifier []type 相比于数组不需要指定大小
		var slice1 []type = make([]type, len)
		slice1 := make([]type, len)
		make([]T, length, capacity) capacity可选，是指定容量，应该是最大值？可以不填
		可以参考数组的定义 helloworld/practice/array.go:21
	3.复制切片、添加切片元素
		append() 和 copy() 函数 这俩函数在builtin.go里面，内置函数吗？
*/
func main() {
	// cap=len=3
	// s :=[] int {1,2,3}
	// s := arr[startIndex:endIndex]
	// s := arr[startIndex:]
	// s := arr[:endIndex]
	// s1 := s[startIndex:endIndex]
	// s :=make([]int,len,cap) // 通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片
	// len() cap()
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var testNil []int
	if testNil == nil {
		fmt.Printf("切片是空的")
	}

	testSliceOpt()

	testAppendAndCopy()
}

func testAppendAndCopy() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func testSliceOpt() {
	/* 创建切片 并初始化 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含) 索引是从0开始算的*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
}

func printSlice(x []int) {
	// len=3 cap=5 slice=[0 0 0]
	fmt.Printf("len=%d cap=%d slice=%d\n", len(x), cap(x), x) // 这里的%d 和 %v是什么区别呢
}
