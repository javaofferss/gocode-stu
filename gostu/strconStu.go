package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
*
strconv 包学习： 包含各种字符串数字类型转换
*/
func main() {
	i, _ := strconv.ParseInt("12", 10, 64)
	fmt.Println(i)                 //打印数字
	fmt.Println(reflect.TypeOf(i)) //打印类型
}
