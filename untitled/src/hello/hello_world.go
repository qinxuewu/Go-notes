
package main //表示一个可独立执行的程序

import "fmt"

/*
	main 函数是每一个可执行程序所必须包含的，
一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
*/
func main() {

	fmt.Println("Hello, World!")
}
