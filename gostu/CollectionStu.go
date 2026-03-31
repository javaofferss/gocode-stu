package main

import "fmt"

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

}

// 测试切片
func testSlice() {
	//类似于java中的ArrayList
	var list = []string{}
	list = append(list, "1")
	list = append(list, "2")
	fmt.Println(list)
}
