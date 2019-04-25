package main   //表示一个可独立执行的程序
import "fmt"

/**
	Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
	声明变量的一般形式是使用 var 关键字：  var identifier type
	 或者根据值自行判定变量类型。  var v_name = value
	 第三种，省略var   	v_name := value
 */
var x, y int
var x1, y2 = 123, "hello"
var (
	// 这种因式分解关键字的写法一般用于声明全局变量
	a1 int
	b1 bool
)


func main() {
	vartest()
}

func vartest() {
	//Go对于已声明但未使用的变量会在编译阶段报错

	//指定变量类型，如果没有初始化，则变量默认为零值
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 第二天中自行判断变量类型
	var name = "qxw"
	fmt.Print("赋值自动匹配变量类型："+name)

	//第三种，省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：

	intVal,intVal2,val3 := 1,2,"字符串3"

	intVal=1111
	fmt.Print(intVal,intVal2,val3)

}

