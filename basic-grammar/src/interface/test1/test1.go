package main

import "fmt"

/**
	interface值

	如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象
 */
type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段
	company string
	money float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("姓名： %s  手机号 %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("员工姓名  %s, 公司  %s. 手机号 %s\n", e.name, e.company, e.phone)
}


// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}



func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("定义Men类型的变量i ,存储Student .........")
	i.SayHi()
	i.Sing("November rain")


	//i也能存储Employee
	i = tom
	fmt.Println("定义Men类型的变量i,i也能存储Employee...............")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("定义了slice Men................")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}

