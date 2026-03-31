package main

import (
	"fmt"
	"reflect"
)

/**
 * 研究 go int
 */
func main() {
	var name = 9                                     //声明一定要用。否则会爆红
	fmt.Println("name的值为：", name)                    //name的值为： 9
	fmt.Println("name值的类型为： ", reflect.TypeOf(name)) //name值的类型为：  int

	var name8 int8 = 9
	fmt.Println("name8 的值为：", name8)
	fmt.Println("name8 的类型为：", reflect.TypeOf(name8)) //name8 的类型为： int8
	name8 = -9
	fmt.Println("name8 的值为：", name8)                  //name8 的值为： -9
	fmt.Println("name8 的类型为：", reflect.TypeOf(name8)) //name8 的类型为： int8
	//name8 = 255; //此处会编译报错：constant 255 overflows int8
	//fmt.Println("name8 的值为：",name8)
	//fmt.Println("name8 的类型为：", reflect.TypeOf(name8))
	/**
	* int 主要有 int, int8 int16 int 32 int64. 无符号对应 uint uint8 uint16 uint32 uint64

	  int 的大小和操作平台有关，由操作平台决定
	*/

}
