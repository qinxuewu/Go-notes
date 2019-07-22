package main

import (
	"fmt"
	"time"
)

// 通道好比就是工厂的传送带, 一个机器（生产者协程）在传送带上放置物品，另外一个机器（消费者协程）拿到物品并打包。
// 未初始化的通道的值是nil。
// ch <- int1 表示：用通道 ch 发送变量 int1
// int2 = <- ch 表示：变量 int2 从通道 ch 接收数据
// 同一个操作符 <- 既用于发送也用于接收，但Go会根据操作对象弄明白该干什么
// 默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束
func main() {
	ch := make(chan string)
	go setData(ch)
	go getData(ch)

	// 让main 协程休眠 防止立马退出
	// 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出
	time.Sleep(1e9)
}

// 用通道 ch 发送变量
func setData(ch chan string) {
	ch <- "aaaaaa"
	ch <- "bbbbbbb"
	ch <- "cccccc"
	ch <- "ccccc"
}

// 接收
func getData(ch chan string) {
	var input string

	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
