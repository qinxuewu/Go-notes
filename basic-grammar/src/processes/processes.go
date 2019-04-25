package main

import (
	"fmt"
)

func main() {
	/*
		Go中流程控制分三大类：条件判断，循环控制和无条件跳转。
	*/
	//var x=10
	//
	//if x>10 {
	//	fmt.Print("大于10")
	//}else{
	//	fmt.Print("小于10")
	//}

	//Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了

	myFor()

}

//Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签
func myFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签
	println(i)
	i++
	goto Here //跳转到Here去
}

//Go里面最强大的一个控制逻辑就是for，它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作
func myFor() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}

	fmt.Println("sum总和= ", sum)

	//另一种写sum2 小于1000 就一直循环 知道满足条件终止
	sum2 := 1
	for sum2 < 10 {
		fmt.Println("sum2= ", sum2)
		sum2 += sum2

	}
	//fmt.Println("sum2= ", sum2)

	//	其中;也可以省略，那么就变成如下的代码了，是不是似曾相识？对，这就是while的功能

	sum3 := 1
	for sum3 < 200 {
		sum3 += sum2
	}
	fmt.Println("sum3= ", sum3)

	//	在循环里面有两个关键操作break和continue	,break操作是跳出当前循环，continue是跳过本次循环
	//	break可以配合标签使用，即跳转至标签所指定的位置

	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Print(index)
	}

	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1

	// break和continue还可以跟着标号，用来跳到多重循环中的外层循环
Href:
	for index := 10; index > 0; index-- {
		if index == 5 {
			break Href
		}
		fmt.Print(index)
	}
	fmt.Printf("%s\n", "    ")

	//	for配合range可以用于读取slice和map的数据：
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s : %s\n", k, v)
	}

	//	可以使用_来丢弃不需要的返回值
	for _, v := range kvs {
		fmt.Printf("%s\n", v)
	}

	// Go里面switch默认相当于每个case最后带有break
	i := 10
	switch i {
	case 1:
		fmt.Printf("%s\n", "i is to 1")
	case 2, 3, 4:
		fmt.Printf("%s\n", "i is to 2,3,4")
	case 10:
		fmt.Printf("%s\n", "i is to 10")
	default:
		fmt.Printf("%s\n", "未匹配到条件 执行默认条件")
	}

	//	使用fallthrough强制执行后面的case代码

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
