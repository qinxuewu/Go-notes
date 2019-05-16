/**
 *
 * @author qinxuewu
 * @create 19/5/16下午10:05
 * @since 1.0.0
 */
package main
import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)


func  Echo(w *webscoker.Conn)  {
	var  err error

	for  {
		var reply string

		if err=websocket.Message.Recevie(ws,&reply); err!=nil{
			fmt.Println("Cat receive")
		}
		fmt.Println("接收到客户端发来的消息 :"+reply)

		msg:="Receive :" +reply
		fmt.Println("发送给客户端的消息 :"+msg)
	}
}

func main()  {
	http.Handle("/",websocket.Handler(Echo))

	if err:=http.ListenAndServe(":1234",nil);err !=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}

