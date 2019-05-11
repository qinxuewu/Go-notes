/**
 *
 * @author qinxuewu
 * @create 19/5/11上午12:17
 * @since 1.0.0
 */
package main

import (
	"net"
	"fmt"
)

//错误检查
func  CheckError(err error)  {
	if err !=nil{
		panic(err)
	}
}

func ProcessInfo(conn net.Conn)  {
	buf:=make([]byte,1024)
	defer  conn.Close()
	for  {
		numOfBytes,err:=conn.Read(buf)
		// 不抛出异常  退出
		if err !=nil{
			break
		}
		if numOfBytes !=0{
			fmt.Printf("messge  is %s\n",string(buf))
		}

	}
}
//go  聊天室 服务端
func main()  {
	listen_scoket,err:=net.Listen("tcp","127.0.0.1:8080")
	CheckError(err)
	defer  listen_scoket.Close()
	fmt.Println("servie start .........")
	for  {
		conn,err:=listen_scoket.Accept()
		CheckError(err)
		go ProcessInfo(conn)
	}
}
