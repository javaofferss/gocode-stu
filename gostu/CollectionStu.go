package main

import (
	"fmt"
	"reflect"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/emirpasic/gods/maps/treemap"
)

func main() {
	testTreeMap()
	testLinkedMap()
	testMap()
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

func testMap() {
	var mapObj map[string]int = make(map[string]int)
	//put
	mapObj["name"] = 1
	fmt.Println(mapObj["name"]) //1
	fmt.Println(mapObj)         //map[name:1]

	//replace
	mapObj["name"] = 2
	fmt.Println(mapObj) //map[name:1]

	//长度
	fmt.Println(len(mapObj)) //1

	//遍历
	for k, v := range mapObj {
		fmt.Println(k, v) //name 2
	}

	//delete
	delete(mapObj, "name")
	fmt.Println(mapObj) //map[]

}

func testTreeMap() {
	comparator := treemap.NewWithStringComparator()
	comparator.Put("c", 1)
	comparator.Put("a", 2)

	for _, key := range comparator.Keys() {
		value, _ := comparator.Get(key)
		fmt.Println(key, value)
		/**
		a 2
		c 1
		*/
	}

	// 这里和java里面api 感觉刚好相反.
	key, value := comparator.Floor("b")
	fmt.Println(key, value) //打印 a 2

	// 这里和java里面api 感觉刚好相反.
	foundKey, foundValue := comparator.Ceiling("b")
	fmt.Println(foundKey, foundValue) //打印 c 1

}

func testLinkedMap() {
	//golang sdk 没有treeMap. 所以需要依赖第三方的. 通过调用NewOrderedMap函数获取
	sortMap := orderedmap.NewOrderedMap[string, int]()
	sortMap.Set("b", 2)
	sortMap.Set("a", 3)
	fmt.Println(sortMap)

	//遍历, 按照插入的顺序遍历
	for key, v := range sortMap.Keys() {
		fmt.Println(key, v)
	}
}
