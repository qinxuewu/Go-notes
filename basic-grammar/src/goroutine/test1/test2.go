/**
 *
 * @author qinxuewu
 * @create 19/7/7下午2:09
 * @since 1.0.0
 */
package main

import (
	"fmt"
	"runtime"
)

func main()  {
	// 使用cpu核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	c :=make(chan  bool,10)  // 创建一个通道 缓存10
	// 运行10个go协程
	for i:=0;i<10;i++ {
		go Go(c,i)
	}

	// 接收10次 
	for i:=0;i<10 ; i++  {
		<-c  //接收任意数据，忽略接收的数据
	}

}

func Go(c chan bool,index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true  //向通道c 发送 true.

}