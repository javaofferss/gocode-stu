package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
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

	//统计字符串的长度
	str := "hello 你好"
	fmt.Println(str+":长度 ", len(str))        //打印长度为12
	fmt.Println(utf8.RuneCountInString(str)) //打印8
	fmt.Println(len([]rune(str)))            //打印8

	//截取字符串
	subStr := str[0:1]
	fmt.Println(subStr) //打印h

	//判断是否包含一次字符串
	c := strings.Contains(str, "hello")
	fmt.Println(c) //打印true

	//大小写转换
	fmt.Println(strings.ToUpper(str))     //HELLO 你好
	fmt.Println(strings.ToLower("HELLO")) //hello

	//字符串转数字
	numStr := "100"
	num, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println(num) // 100

	numStr = "+200"
	fmt.Println(strconv.ParseInt(numStr, 10, 64))

	numStr = "-200"
	fmt.Println(strconv.ParseInt(numStr, 10, 64))

}
