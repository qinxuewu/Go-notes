
/**
	嵌入interface

	Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，
	那么相同的逻辑引入到interface里面，那不是更加完美了。
	如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。


 */
package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}


//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("姓名： %s  手机号 %s\n", h.name, h.phone)
}

type Element1 interface {
	SayHi()
}

type Element2 interface {
	Element1
}

func main()  {
	h:=Human{"Mike", 25, "222-222-XXX"}

	//定义Men类型的变量i
	var i Element1

	//i能存储Student
	i = h
	i.SayHi()

	var e Element2
	e=h
	e.SayHi()
}