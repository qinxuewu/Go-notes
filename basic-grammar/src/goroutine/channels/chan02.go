package main

import "fmt"

// 默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
// 这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	// 创建无缓冲的通道
	c := make(chan int)
	go sum2(s[:len(s)/2], c) //取前3个  [7 2 8]
	go sum2(s[len(s)/2:], c) // 跳过前3个 -9, 4, 0
	x := <-c                 // 从 c 中接收
	y := <-c
	fmt.Println(x, y)
}

func sum2(s []int, c chan int) {
	count := 0
	for _, v := range s {
		count = count + v
	}
	// 将count发送到通道中  “箭头”就是数据流的方向
	c <- count
}
