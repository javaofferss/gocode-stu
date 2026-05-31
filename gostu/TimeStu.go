package main

import (
	"fmt"
	"time"
)

/**
 * time stu
 */
func main() {
	//今日日期
	now := time.Now()
	fmt.Println(now)

	//打印年月日
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Weekday())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	//格式化
	fmt.Println(fmt.Sprintf("%02d-%02d-%02d", now.Year(), now.Month(), now.Day()))

	//格式化: 和java的格式化不一样。并不是传统的Y-M-d HH:mm:SS.
	//而是：2006 表示年。01： 表示月。02：表示天。15表示时。 04表示分。05表示秒
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//把字符格式化成日期
	str := "2026-06-31"
	fmt.Println(time.Parse(str, "2006-01-02"))

	//打印毫秒
	fmt.Println(now.UnixMilli())
}
