package main

/*
*

	if for
*/
func main() {
	var money = 1000
	if money == 1000 {
		println("money is 1000")
	}

	if money > 1000 {
		println("money 大于 1000")
	}

	for i := 0; i < money; i = money {
		println("这是 for 循环，我在表达式中没有 （ ）， 我有三个模块")
	}
	for i := 0; i < money; {
		println("这是 for 循环，我的表达式中没有 （）， 我有两个模块")
		i = money
	}
	for i := 0; ; {
		println("我只有一个模块，我利用break 跳出循环，go 语言要求声明的变量必须要用。所以我打印了i:", i)
		break //
	}
	for {
		println("我一个模块都没有。但是我利用break,可以跳出去")
		break
	}

}
