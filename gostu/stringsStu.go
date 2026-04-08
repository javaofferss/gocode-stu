package main

import (
	"fmt"
	"strings"
)

func main() {
	//这个count 不是单独的统计字符的长度。他有一个特性。可以看一眼源码。
	count := strings.Count("你好", "")
	fmt.Println(count)

	//打印字符b在 字符串abc中的下标位置
	var index = strings.Index("abc", "b")
	fmt.Println(index)
	index = strings.Index("abc", "a")
	fmt.Println(index)
}
