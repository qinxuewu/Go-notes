package main

import (
	"fmt"
	"time"
)

/**
	如果存在多个channel的时候，我们该如何操作呢，Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。

	select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
 */
func main()  {
	c:=make(chan  int)
	quit:=make(chan  int)

	go func() {
		for i:=0;i<10 ;i++  {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)
}

/**
 定义两个通道c,quit 都是int类型
 */
func  fib(c,quit chan  int)  {
	x,y:=1,1

	//循环接收
	for  {
		select {
			//发送x 到channel c.
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
	
}

//有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时
func timetest()  {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <- c:
				println(v)
			case <- time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<- o
}