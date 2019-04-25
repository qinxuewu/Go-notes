package main

import "fmt"

/**
Go语言中，也和C或者其他语言一样，我们可以声明新的类型,作为其它类型的属性或字段的容器。
例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。
*/
type person struct {
	name string
	age  int
}

func main() {
	var P person // P现在就是person类型的变量了

	//赋值操作
	P.name = "qxw"
	P.age = 16
	fmt.Printf("The person's name is %s\n", P.name) // 访问P的name属性.

	//	另外几种声明使用方式

	//  1.按照顺序提供初始化值
	P2 := person{"qqqq", 11}
	fmt.Printf("按照顺序提供初始化值: %v\n", P2.name)

	//	2.通过field:value的方式初始化，这样可以任意顺序

	P3 := person{age: 22, name: "dddddd"}
	fmt.Printf("通过field:value的方式初始化，这样可以任意顺序: %v\n", P3.name)

	// 3.当然也可以通过new函数分配一个指针，此处P的类型为*person
	P4 := new(person)

	fmt.Printf("通过new函数分配一个指针，此处P的类型为*person: %v\n", P4.name)

	var tom person
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{"bob", 20}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 30}

	tb_Older, tb_diff := Older(tom, bob)
	fmt.Printf("%s,%s,%d,%v\n ", tom.name, bob.name, tb_diff, tb_Older)

	tb_Older2, tb_diff2 := Older(tom, paul)
	fmt.Printf("%s,%s,%d,%v\n ", tom.name, paul.name, tb_diff2, tb_Older2)

	//structTest()

	structTest2()

}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

/**
struct的匿名字段
我们上面介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，
实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

类似JAVA对象中声明其另一个对象作为其对象字段
*/

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func structTest() {
	// 我们初始化一个学生
	mark := Student{Human{"aa", 11, 20}, "testetetet"}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("His speciality is ", mark.speciality)

	// 修改他的年龄信息
	mark.age = 46
	fmt.Println("His age is", mark.age)

	// 修改他的体重信息
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	//	student还能访问Human这个字段作为字段名
	mark.Human.age = 1
}

//  不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段
type Skills []string

type Human2 struct {
	name   string
	age    int
	weight int
}

type Student2 struct {
	Human2     // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func structTest2() {
	jane := Student2{Human2: Human2{"a", 1, 1}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)

	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

type Human3 struct {
	name  string
	age   int
	phone string // Human类型拥有的字段
}

type Employee struct {
	Human3     // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func structTest3() {
	Bob := Employee{Human3{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human3.phone)
}
