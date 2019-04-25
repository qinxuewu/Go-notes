package main

import (
	"fmt"
	"math"
)

/**
   面向对象
 */
func  main()  {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))


	methodTest()
}

/**
	现在假设有这么一个场景，你定义了一个struct叫做长方形，你现在想要计算他的面积，
	那么按照我们一般的思路应该会用下面的方式来实现


 */
type Rectangle struct {
	width, height float64
}

/**
	这段代码可以计算出来长方形的面积，但是area()不是作为Rectangle的方法实现的（类似面向对象里面的方法），
	而是将Rectangle的对象（如r1,r2）作为参数传入函数计算面积的。
	这样实现当然没有问题咯，但是当需要增加圆形、正方形、五边形甚至其它多边形的时候,就麻烦了 要写多个计算方法
 */
func area(r Rectangle) float64 {
	return r.width*r.height
}

/**
	用method来实现
 */

type Rectangle2 struct {
	width, height float64
}

type Circle struct {
	radius float64
}

/**
	类似java里的方法重载 根据传入不同的对象 选择对应的计算方法
	此处方法的Receiver是以值传递，

	Receiver还可以是指针, 两者的差别在于,
			指针作为Receiver会对实例对象的内容发生操作,
			而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作
 */
func (r Rectangle2) area2() float64 {
	return r.width*r.height
}

func (c Circle) area2() float64 {
	return c.radius * c.radius * math.Pi
}

/**
	在使用method的时候重要注意几点

	虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
	method里面可以访问接收者的字段
	调用method通过.访问，就像struct里面访问字段一样
 */
func methodTest() {
	r1 := Rectangle2{12, 2}
	r2 := Rectangle2{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area2())
	fmt.Println("Area of r2 is: ", r2.area2())
	fmt.Println("Area of c1 is: ", c1.area2())
	fmt.Println("Area of c2 is: ", c2.area2())
}


