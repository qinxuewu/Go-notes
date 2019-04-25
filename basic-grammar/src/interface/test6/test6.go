package main

import (
	"reflect"
	"fmt"
)

/**
	Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。

使用reflect一般分成三步，下面简要的讲解一下：要去反射是一个类型的值(这些值都实现了空interface)，
首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)
 */

func main()  {
	//test1()

	test2()
}

func test1()  {
	//获取反射值能返回相应的类型和数值
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	
}

//如果要修改相应的值，必须这样写
func test2()  {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	fmt.Println("value:", v.Float())
	v.SetFloat(7.1)
	fmt.Println("value:", v.Float())


}