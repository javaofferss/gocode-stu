package main

import (
	"fmt"
	"reflect"
)

func main() {
	testSlice()
	testSet()

}

// 测试集合
func testSet() {
	//golang中没有提供原生的Set。可以使用Map代替
	var set = make(map[string]bool)
	set["1"] = true
	set["2"] = true
	set["1"] = true
	fmt.Println(set)
	fmt.Println(set["1"]) //这里可以看作包含. 返回true
	fmt.Println(set["4"]) // 这里不存在，返回false
	delete(set, "1")      //删除key
	fmt.Println(set)

}

// 测试切片
func testSlice() {
	//类似于java中的ArrayList
	var list = []string{}
	list = append(list, "1")
	list = append(list, "2")
	fmt.Println(list)

	//测试截取
	fmt.Println(list[0:1]) //[1]
	fmt.Println(list[0:0]) //[]
	fmt.Println(list[1:])  //[2]

	//长度
	fmt.Println(len(list)) //2

	//删除元素
	newList := list[1:]
	fmt.Println(newList, reflect.TypeOf(newList)) //[2] []string
	newList3 := append(newList, "3")
	var addr = &newList
	fmt.Println(addr, reflect.TypeOf(addr)) //&[2] *[]string
	fmt.Printf("%p\n", addr)                //0x72e4780ba060
	fmt.Println(newList3, &newList3)        //[2 3] &[2 3]
}
