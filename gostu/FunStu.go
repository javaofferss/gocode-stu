package main

import "fmt"

// 定义了一个函数类型.
type PrintFun func(string)

func outPrint(str string) PrintFun {
	//返回一个函数类型
	return func(a string) {
		fmt.Println(str)
		fmt.Println(a)
	}
}

func main() {
	fun := outPrint("hello world")
	fun("hello world a")

	/**
	hello world
	hello world a
	*/
}
