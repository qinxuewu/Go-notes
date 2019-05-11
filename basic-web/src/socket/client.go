/**
 *  客户端
 * @author qinxuewu
 * @create 19/5/11上午12:28
 * @since 1.0.0
 */
package main

import (
	"net"
	"fmt"
)

func main()  {
		connn,err:=net.Dial("tcp","127.0.0.1:8080")
		if err !=nil{
			panic(err)
		}
		defer  connn.Close()

		connn.Write([]byte("hello  go  start 22大萨达撒多"))

		fmt.Println("clien .............")
}
