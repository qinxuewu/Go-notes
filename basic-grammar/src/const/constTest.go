package main //表示一个可独立执行的程序
import "fmt"

/**
  常量
	所谓常量，也就是在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。
	在Go程序中，常量可定义为数值、布尔值或字符串等类型。
 */

 //自动匹配类型
const age  =25
//指定类型
const name string ="qxwxwxx"
//另一种方法
const(
	age1=26
	name2="ddsdsdsds"
)

//内置基础类型
var isActive bool  // 全局变量声明
var isFalse,iseTrue = false,true   // 忽略类型的声明


//Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，const中每增加一行加1：
const(
	f1=iota  //f1=0
	f2=iota  //f2=1
	f3=iota   //f3=2
	f4	//常量声明省略值时，默认和之前一个值的字面相同,隐式地说f4 = iota , 因此 f4=3
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

func main() {

	const Pi = 3.1415926
	const i = 10000
	const MaxThread = 10
	const prefix = "astaxie_"



	fmt.Println(Pi*i)


	fmt.Println(name,age,name2,age1)

	//在Go中字符串是不可变的
	var s string = "hello"
	//s[0] = 'c'   // 代码编译时会报错：cannot assign to s[0]

	fmt.Println(s)

	//但如果真的想要修改怎么办呢？下面的代码可以实现
	c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c)  // 再转换回 string 类型
	fmt.Printf("%s\n", s2)


	//修改字符串也可写为：
	s = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s)



	//Go中可以使用+操作符来连接两个字符串：
	mm := "初学者"
	cc := "开始学go"
	result :=mm+cc
	fmt.Printf("%s\n",result)



	/**
		格式化输入输出
		v：默认格式，不同类型的默认格式如下：

		　　布尔型：t
		　　整　型：d
		　　浮点型：g
		　　复数型：g
		　　字符串：s
		　　通　道：p
		　　指　针：p
	 */
	//iota
	fmt.Print(f1,f2,f3,f4)

	fmt.Printf("%s\n","")

	//每遇到一个const关键字，iota就会重置，此时v == 0
	fmt.Printf("%d\n",v)
}





