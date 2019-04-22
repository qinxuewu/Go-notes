package main

import "fmt"

/**

	method继承
	method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
 */

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	//  包含这个匿名字段的struct也能调用该method
	mark.SayHi()
	sam.SayHi()
}

