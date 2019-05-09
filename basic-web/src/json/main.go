/**
 *  解析JSON
 * @author qinxuewu
 * @create 19/5/9下午9:46
 * @since 1.0.0
 */
package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	test()
	test3()
	test5()
}

//SON输出的时候必须注意，只有导出的字段才会被输出，如果修改字段名，那么就会发现什么都不会输出，所以必须通过struct tag定义来实现
//  创建结构体 解析json
type Server struct {
		ServerName string `json:"serverName"`
		ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

//JSON解析的时候只会解析能找得到的字段，找不到的字段会被忽略
func test()  {
	var s Serverslice
	// 待解析的json字符串
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s)
}

//解析到interface 未知json格式
func test3(){
	//假设有如下的JSON数据
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err :=json.Unmarshal(b,&f)
	if err !=nil {
		panic(err)
	}

	//这个时候f里面存储了一个map类型，他们的key是string，值存储在空的interface{}里
	fmt.Println(f)

	//访问这些数据
	m :=f.(map[string]interface{})
	age:=m["Age"]
	fmt.Println("age",age)

	//断言访问
	for k,v :=range  m{
		switch vv:=v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i,u:=range vv{
				  fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}

	}
}



//生成JSON

func test5()  {
	var s Serverslice
		s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
     	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
    	b, err := json.Marshal(s)
    	if err != nil {
    		fmt.Println("json err:", err)
    	}
    	fmt.Println("生成json",string(b))
}











