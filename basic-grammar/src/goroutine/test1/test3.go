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
	"sync"
)

func main()  {
	// 使用cpu核数
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg :=sync.WaitGroup{}
	wg.Add(10)

	for i:=0;i<10;i++ {
		go Go3(&wg,i)
	}
	// 等待程序执行完毕
	wg.Wait()


}

func Go3(wg *sync.WaitGroup,index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	wg.Done()

}