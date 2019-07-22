package main

import (
	"fmt"
	"time"
)

// Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64
var week time.Duration

// 时间和日期
func main() {
	t := time.Now()
	fmt.Println("获取当期时间:", t)

	//	自定义时间格式化字符串
	fmt.Printf("时间格式化:  %02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	// UTC 表示通用协调世界时间。
	t = time.Now().UTC()
	fmt.Println("UTC 表示通用协调世界时间: ", t)

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println("week_from_now:", week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
