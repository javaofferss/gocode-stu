package main

import (
	"fmt"
	"reflect"
)

/**
 * 接口 可以理解为任意类型.
 */
type anyObj interface{}

func GetType(obj anyObj) string {
	switch obj.(type) {
	case string:
		//转成string类型
		s := obj.(string)
		return "string" + s
	case int:
		return "int"
	case float64:
		return "float64"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

type OrderType int

const (
	Phone OrderType = iota
	Book
)

type Order interface {
	GetOrderName() string
	GetOrderType() OrderType
}

// 接口不能有方法实现 - 接口只定义方法签名。所以不能这么写
//func (order Order) printOrderInfo() {
//	orderName := order.GetOrderName()
//	fmt.Println(orderName)
//	orderType := order.GetOrderType()
//	fmt.Println(orderType)
//}

func printOrderInfo(order Order) {
	orderName := order.GetOrderName()
	fmt.Println("订单名称：", orderName)
	orderType := order.GetOrderType()
	fmt.Print("订单类型：")
	switch orderType {
	case Phone:
		fmt.Println("Phone")
	case Book:
		fmt.Println("Book")
	default:
		fmt.Println("unknown")
	}
}

type PhoneOrder struct {
	name string
}

func (p *PhoneOrder) GetOrderName() string {
	return p.name
}

func (p *PhoneOrder) GetOrderType() OrderType {
	return Phone
}

type BookOrder struct {
	name string
}

func (b *BookOrder) GetOrderName() string {
	return b.name
}

func (b *BookOrder) GetOrderType() OrderType {
	return Book
}

func main() {
	//打印类型
	var str string = "hello"
	var any anyObj = str
	fmt.Println(any, reflect.TypeOf(any))

	//通过switch 来匹配类型
	fmt.Println(GetType(any))

	//接口可做模板
	phoneOrder := &PhoneOrder{
		name: "HuaWei",
	}

	//转成Order对象。饭后 再调用
	var order Order = phoneOrder
	printOrderInfo(order)

	bookOrder := &BookOrder{
		name: "golang 学习指南",
	}

	//直接调用Order方法
	printOrderInfo(bookOrder)

}
