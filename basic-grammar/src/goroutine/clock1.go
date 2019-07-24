package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//
func main() {
	// 监听一个网络端口上到来的连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Accept方法会直接阻塞，直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// go关键字支持多客户端连接，让每一次handleConn的调用都进入一个独立的goroutine
		go handleConn(conn)
		//handleConn(conn)
	}

}

// 处理一个完整的客户端连接
func handleConn(c net.Conn) {
	fmt.Println("客户端链接加入：", c.RemoteAddr())
	defer c.Close()
	// 死循环
	for {
		// 获取当前时刻，然后写到客户端
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))

		if err != nil {
			return
		}
		// 每个已秒执行一次  知道客户端断开连接
		time.Sleep(1 * time.Second)
	}
}
