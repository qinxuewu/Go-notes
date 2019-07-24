package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		keyValue := strings.Split(v, "=")
		go connTcp(keyValue[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

// 获取输入的端口号建立连接
func connTcp(uri string) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy2(os.Stdout, conn)

}

func mustCopy2(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
