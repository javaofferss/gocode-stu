package main

import "reflect"

func main() {
	var money float32 = 100.000000001
	println("money is ", money)
	println("type of ", reflect.TypeOf(money))

	var money64 float64 = 100.000000001
	println("money64 is : ", money64)
	println("type of ", reflect.TypeOf(money64))

	/** 打印结果：
	money is  +1.000000e+002
	type of  (0x4b9b80,0x484900)
	money64 is :  +1.000000e+002
	type of  (0x4b9b80,0x484940)

	*/

}
