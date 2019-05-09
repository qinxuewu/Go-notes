package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
func main()  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	//-------- 字符串类型操作  ----------------------------------------------------------

	// 写入值
	_, err = c.Do("SET", "key_name", "秦学武")
	checkErr(err)

	//获取值 get
	name,err:= redis.Strings(c.Do("GET","key_name"))
	checkErr(err)
	fmt.Println("key_name: ",name)


	

}


func checkErr(err error)  {
	if err !=nil{
		panic(err)
	}
}