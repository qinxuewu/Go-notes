package main

import (
	"fmt"
	"math"
)

/**
面向对象  method

	前面两章我们介绍了函数和struct，那你是否想过函数当作struct的字段一样来处理呢？
	今天我们就讲解一下函数的另一种形态，带有接收者的函数，我们称为method

*/
func main() {
	r1 := Rectanlge{12, 2}
	r2 := Rectanlge{9, 4}
	fmt.Println("Area of r1 is: ", area(r1))
	fmt.Println("Area of r2 is: ", area(r2))

	methodTest()
}

//现在假设有这么一个场景，你定义了一个struct叫做长方形，你现在想要计算他的面积，那么按照我们一般的思路应该会用下面的方式来实现
type Rectanlge struct {
	width, height float64
}

//计算面积
func area(r Rectanlge) float64 {
	return r.width * r.height

}

/**
method的语法实现   func (r ReceiverType) funcName(parameters) (results)


*/

type Circle struct {
	radius float64
}

func (r Rectanlge) area22() float64 {
	return r.width * r.height
}

func (c Circle) area22() float64 {
	return c.radius * c.radius * math.Pi
}

/**
在使用method的时候重要注意几点
	虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
    method里面可以访问接收者的字段
	调用method通过.访问，就像struct里面访问字段一样
*/
func methodTest() {
	r1 := Rectanlge{3, 6}
	r2 := Rectanlge{3, 8}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area22())
	fmt.Println("Area of r2 is: ", r2.area22())
	fmt.Println("Area of c1 is: ", c1.area22())
	fmt.Println("Area of c2 is: ", c2.area22())

}

/**

你可以在任何的自定义类型中定义任意多的method，接下来让我们看一个复杂一点的例子
*/

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte // Color作为byte的别名

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // Box数组

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func methodTest2() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
