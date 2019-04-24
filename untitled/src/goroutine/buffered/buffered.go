package main

import "fmt"

/**
Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel

在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
*/
func main() {

	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)

	//修改为1报如下的错误:
	//fatal error: all goroutines are asleep - deadlock!

	fmt.Println("-----------------------分割线-----------------------------")

	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}
}

/**
上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，所以也可以通过range，像操作slice或者map一样操作缓存类型的channel
*/
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
