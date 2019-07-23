package main

import "fmt"

// 在Go语言中，你可以给任意类型（包括内置类型，但不包括指针类型）

type Integer int

func (a Integer ) Less(b Integer) bool  {   // 面向对象
	return  a < b
}

func Integer_Less(a Integer, b Integer) bool { // 面向过程
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
func  main()  {
	var a Integer =1
	if a.Less(2) {
		// a小于B 则输出
		fmt.Println(a,"Less 2")
	}

	 b:=Integer_Less(a, 2) // 面向过程的用
	 fmt.Println(b)

	 // 指针操作
	a.Add(2)
	fmt.Println("a =", a)
}
