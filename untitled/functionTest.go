package main

import "fmt"

/*
	函数是Go里面的核心设计，它通过关键字func来声明  它的格式如下：

	func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
		//这里是处理逻辑代码
		//返回多个值
		return value1, value2
	}
*/
func main() {
	x := 3
	y := 4
	z := 5
	max1 := max(x, y)
	max2 := max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max1)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max2)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它

	sum1, sum2 := sumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, sum1)
	fmt.Printf("%d * %d = %d\n", x, y, sum2)

}

/**
	关键字func用来声明一个函数funcName
    函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
	函数可以返回多个值
	上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
	如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
	如果没有返回值，那么就直接省略最后的返回信息
	如果有返回值， 那么必须在函数的外层添加return语句
*/

//max函数有两个参数，它们的类型都是int，那么第一个变量的类型可以省略（即 a,b int,而非 a int, b int)，默认为离它最近的类型
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Go语言比C更先进的特性，其中一点就是函数能够返回多个值。

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

//Go函数支持变参。接受变参的函数是有着不定数量的参数的
//arg ...int告诉Go这个函数接受不定数量的参数。
// 注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice：
func myFunc2(arg ...int) {
	for _, n := range arg {
		fmt.Printf("变参函数打印: %d\n", n)
	}
}
