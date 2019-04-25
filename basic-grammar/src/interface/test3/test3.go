package main

import (
	"strconv"
	"fmt"
)

/**
	interface函数参数

	interface的变量可以持有任意实现该interface类型的对象

	我们是不是可以通过定义interface参数，让函数接受各种类型的参数。
 */

type Human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}


func main()  {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
