package main

import (
	"fmt"
)

// 通过一个（或多个）通道交换数据进行协程同步
func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	//out := make(chan int, 1) // solution 2
	out <- 2
	go f1(out)
}