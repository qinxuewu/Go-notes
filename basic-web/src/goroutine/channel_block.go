package main

import (
	"fmt"
	"time"
)

// 对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的
// 对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用
func main() {
	ch1 := make(chan int)
	go pump(ch1)
	go suck(ch1)
	//fmt.Println(<-ch1)
	time.Sleep(1e9)
}

// 一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字 0。
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 为通道解除阻塞定义了 suck 函数来在无限循环中读取通道
func suck(ch chan int) {

	for {
		fmt.Println(<-ch)
	}
}
