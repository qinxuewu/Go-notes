package main

import "fmt"

/**
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。那么goroutine之间如何进行数据的通信呢
Go提供了一个很好的通信机制channel。channel可以与Unix shell 中的双向管道做类比:可以通过它发送或者接收值。
这些值只能是特定的类型：channel类型。 定义一个channel时，也需要定义发送到channel的值的类型。注意，必须使用make 创建channel：
*/

func main() {
	a := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sum(a[:len(a)/2], c) //取值范围 前两个 1,2
	go sum(a[len(a)/2:], c) //取值范围，从下标2开始取值  3,4,5
	x := <-c
	y := <-c
	fmt.Println(x, y, x+y)

	//		默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock

}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total = total + v
	}

	/**
	  	channel通过操作符<-来接收和发送数据

	      ch <- v    // 发送v到channel ch.
	  	v := <-ch  // 从ch中接收数据，并赋值给v
	*/
	c <- total //发送total  到channel c.
}
