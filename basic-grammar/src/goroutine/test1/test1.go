/**
 *
 * @author qinxuewu
 * @create 19/7/7下午2:09
 * @since 1.0.0
 */
package main

import "fmt"

func main()  {
	c:=make(chan  bool)
	go func() {
		fmt.Println("GO GO !!")
			c <- true
			close(c)
	}()

	for v:=range c {
		fmt.Println(v)
	}
	// 输出 go go
}

