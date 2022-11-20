package main

import "fmt"

/*
*

	map数据结构
	1.key value 无序键值对的格式
	2.定义
		// 声明变量，默认 map 是 nil
		// var map_variable map[key_data_type]value_data_type

		// 使用 make 函数
		//map_variable := make(map[key_data_type]value_data_type)
	3.遍历 for key := range map
	4.delete函数
*/
func main() {
	var countryCapitalMap map[string]string /*创建集合 */

	if countryCapitalMap == nil {
		println("这里只是声明，没有初始化，所以集合为空，nil map 不能用来存放键值对")
	}

	countryCapitalMap = make(map[string]string)

	if countryCapitalMap != nil {
		println("这里把声明进行了初始化，所以集合已经不为空了")
		println("countryCapitalMap:", countryCapitalMap)
	}

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 value exist*/
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不在map中")
	}

	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")
	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}

// 基于go的简单hashMap，未做key值校验
/*package main

import (
"fmt"
)

type HashMap struct {
	key string
	value string
	hashCode int
	next *HashMap
}

var table [16](*HashMap)

func initTable() {
	for i := range table{
		table[i] = &HashMap{"","",i,nil}
	}
}

func getInstance() [16](*HashMap){
	if(table[0] == nil){
		initTable()
	}
	return table
}

func genHashCode(k string) int{
	if len(k) == 0{
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		hashCode += (hashCode + int(k[i]))*31
	}
	return hashCode
}

func indexTable(hashCode int) int{
	return hashCode%16
}

func indexNode(hashCode int) int {
	return hashCode>>4
}

func put(k string, v string) string {
	var hashCode = genHashCode(k)
	var thisNode = HashMap{k,v,hashCode,nil}

	var tableIndex = indexTable(hashCode)
	var nodeIndex = indexNode(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var headNode = headPtr[tableIndex]

	if (*headNode).key == "" {
		*headNode = thisNode
		return ""
	}

	var lastNode *HashMap = headNode
	var nextNode *HashMap = (*headNode).next

	for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex){
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue string = lastNode.value
		lastNode.value = thisNode.value
		return oldValue
	}
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func get(k string) string {
	var hashCode = genHashCode(k)
	var tableIndex = indexTable(hashCode)

	var headPtr [16](*HashMap) = getInstance()
	var node *HashMap = headPtr[tableIndex]

	if (*node).key == k{
		return (*node).value
	}

	for (*node).next != nil {
		if k == (*node).key {
			return (*node).value
		}
		node = (*node).next
	}
	return ""
}

//examples
func main() {
	getInstance()
	put("a","a_put")
	put("b","b_put")
	fmt.Println(get("a"))
	fmt.Println(get("b"))
	put("p","p_put")
	fmt.Println(get("p"))
}*/
