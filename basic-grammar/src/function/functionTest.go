package main

import (
	"fmt"
	"os"
)

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

	//	传值与传指针
	a := 3
	fmt.Printf("%d\n", a) // 顺序胡 x=3
	a1 := myFunc3(a)      //调用函数

	fmt.Println("a+1 = ", a1) // 应该输出"x+1 = 4"
	fmt.Println("a= ", x)     // 应该输出"x = 3"

	//	指针修改参数
	a2 := myFUnc4(&a)

	fmt.Printf("指针操作-》 x+1 =  %d\n", a2)   // 应该输出 "x+1 = 4"
	fmt.Printf("原始参数a也修改-》 x =   %d\n", a) // 应该输出 "x = 4"

	mydefer()
	mydefer2()

	fmt.Printf("%s\n", "   ")

	//	函数作为值、类型

	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Printf("初始化-> slice =  %v\n", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递了
	fmt.Printf("odd: %v\n: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递了
	fmt.Printf("isEven:   %v\n ", even)

	//panicTest()
	falg := throwsPanic(panicTest)
	fmt.Printf("让进入panic状态的goroutine恢复过来,recover仅在延迟函数中有效    %t\n", falg)
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

/*
	传值与传指针
	当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，
	当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上
*/
func myFunc3(a int) int {
	a = a + 1 //改变a的值
	return a  //返回新值
}

/**
	指针操作，copy传入参数的指针，这样可以修改传入变量的值

 	传指针使得多个函数能操作同一个对象。
	传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。

	如果用参数值传递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。
	所以当你要传递大的结构体的时候，用指针是一个明智的选择。

	Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针
	（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/
func myFUnc4(a *int) int {
	*a = *a + 1
	return *a
}

/*
	Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。
	当函数执行到最后时，这些defer语句会按照逆序执行，最后该函数返回。
	特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。
*/
func mydefer() {
	defer fmt.Printf("%s\n", "这里defer语句会在退出前调用 输出这句话")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", i)
	}
}

//	如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
func mydefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

/**
函数作为值、类型
在Go中函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
*/

// 声明一个函数类型

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer&2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数
// 函数当做值和类型在我们写一些通用接口的时候非常有用

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		//通过函数类型 判断这个值返回是否为true
		if f(value) {
			//如果是true则追加的数组中
			result = append(result, value)
		}
	}
	return result
}

/**
	Panic和Recover

	Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制
	一定要记住，你应当把它作为最后的手段来使用，也就是说，你的代码中应当没有，或者很少有panic的东西。

	Panic：是一个内建函数，可以中断原有的控制流程，进入一个panic状态中。
           当函数F调用panic，函数F的执行被中断，但是F中的延迟函数会正常执行，然后F返回到调用它的地方。
           在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到发生panic的goroutine中所有调用的函数返回，此时程序退出。
           panic可以直接调用panic产生。也可以由运行时错误产生，例如访问越界的数组。

	Recover：是一个内建的函数，可以让进入panic状态的goroutine恢复过来。
			recover仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil，并且没有其它任何效果。
			如果当前的goroutine陷入panic状态，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
*/
var user = os.Getenv("USER")

func panicTest() {
	if user == "" {
		panic("no value for $USER  中断原有的控制流程，进入一个panic状态中")
	}

}

//这个函数检查作为其参数的函数在执行时是否会产生panic：
func throwsPanic(f func()) (b bool) {
	//延迟语句
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}

/**
main函数和init函数
Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）
这两个函数在定义时不能有任何的参数和返回值
虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，
我们都强烈建议用户在一个package中每个文件只写一个init函数。


Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数
每个package中的init函数都是可选的，但package main就必须包含一个main函数。
当一个包被导入时，如果该包还导入了其它的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话）

等所有被导入的包都加载完毕了，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在的话）

点操作
	import(
	 . "fmt"
	)
	这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名 可以省略的写成Println("hello world")



别名操作
	别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
	 import(
	 f "fmt"
	)
	即f.Println("hello world")

_操作
	这个操作经常是让很多人费解的一个操作符
		import (
			"database/sql"
			_ "github.com/ziutek/mymysql/godrv"
		)
	_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
*/
