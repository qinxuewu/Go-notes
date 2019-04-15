package main

import "fmt"

func main() {

	/*
		map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int｀类型，
		而map多了很多类型，可以是int，可以是string及所有完全定义了==与!=操作的类型。
	*/

	//// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var map1 map[string]int

	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map1["two"])

	/*
		使用map过程中需要注意的几点：
			map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
			map的长度是不固定的，也就是和slice一样，也是一种引用类型
			内置的len函数同样适用于map，返回map拥有的key的数量
			map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
			map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
	*/

	fmt.Printf("%d\n", len(map1))

	//通过delete删除map的元素

	// 初始化一个字典
	map2 := map[string]float32{"a": 1.1, "b": 2.1, "c": 3.1, "d": 4}
	fmt.Printf("%v\n", map2)

	delete(map2, "a") // 删除key为C的元素
	fmt.Printf("删除k为a的元素===   %v\n", map2)

	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := map2["a"]
	if ok {
		fmt.Println("指定的k存在 ", csharpRating)
	} else {
		fmt.Println("指定的k不存在")
	}

	//map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了

	fmt.Printf("%v\n", m)

	/*
			make、new操作

			make用于内建类型（map、slice 和channel）的内存分配。
		    new用于各种类型的内存分配。
	*/
}
