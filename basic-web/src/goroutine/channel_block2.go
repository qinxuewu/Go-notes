package main

import (
	"fmt"
	"time"
)

// 写一个通道证明它的阻塞性，开启一个协程接收通道的数据，持续 15 秒，然后给通道放入一个值。在不同的阶段打印消息并观察输出
func main() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("接收: ", x)
	}()

	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
