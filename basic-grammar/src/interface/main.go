package main

import "fmt"

/**
	Go语言里面设计最精妙的应该算interface，它让面向对象，内容组织实现非常的方便
	，interface是一组method签名的组合，我们通过interface来定义对象的一组行为

	interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口
 */

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human		//匿名字段Human
	school string
	loan float32
}

type Employee struct {
	Human //匿名字段Human
	company string
	money float32
}

//Human对象实现Sayhi方法
func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}



/**
	interface可以被任意的对象实现。Men interface被Human、Student和Employee实现。
	一个对象可以实现任意多个interface。上面的Student实现了Men和YoungChap两个interface。
 */

// 定义interface

type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}
func main()  {
	
}
