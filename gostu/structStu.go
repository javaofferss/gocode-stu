package main

import "fmt"

type Persion struct {
	Man
}

type Man struct {
	name string
	age  int32
}

// 对象一定要传递指针.
func (p *Persion) setName(name string) {
	p.name = name
}

// 设置name无效. 因为是值传递。非指针.
func (p Persion) setNameInvalid(name string) {
	p.name = name
	//这里会打印p的信息被修改了，但是在调完方法外，依然是原来的值
	fmt.Println(fmt.Sprint("====> ", p))
}

func main() {
	var p = Persion{
		Man{
			name: "cmj",
			age:  22,
		},
	}
	fmt.Println(p)

	p.name = "zyj"
	fmt.Println(p)

	p.setNameInvalid("hello")
	fmt.Println(p)

	p.setName("good")
	fmt.Println(p)

}
