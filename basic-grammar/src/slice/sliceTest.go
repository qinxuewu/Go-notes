package main

import "fmt"

func main() {

	/**
	在很多应用场景中，数组并不能满足我们的需求。在初始定义数组时，我们并不知道需要多大的数组
	因此我们就需要“动态数组”。在Go里面这种数据结构叫slice

	slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。
	*/

	//语法： var fslice []int   和声明array一样，只是少了长度

	sliceArr := []byte{'a', 'b', 'c', 'd'}

	fmt.Printf("%v\n", sliceArr)

	/**
	slice可以从一个数组或一个已经存在的slice中再次声明。
	slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，但不包含array[j]，它的长度是j-i。
	*/

	// 声明一个含有10个元素元素类型为byte的数组
	var arr1 = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fmt.Printf("arr1== %v\n", arr1)

	//声明两个slice类型的byte
	var a, b []byte

	//赋值操作   a指向 arr1数组的第3个元素开始，并到第五个元素结束，
	a = arr1[2:5]
	fmt.Printf("a== %v\n", a)

	// b指向 arr1数组的第1个元素开始，并到第2个元素结束，
	b = arr1[1:2]
	fmt.Printf("b==  %v\n", b)

	//slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符

	sliceArr2 := []int{1, 2, 4, 5, 6}
	fmt.Printf("sliceArr2== %v\n", sliceArr2)

	var c, d []int

	c = sliceArr2[:]   //  指向 sliceArr2数组
	d = sliceArr2[1:3] //  指向 sliceArr2数组的第2个元素开始，并到第3个元素结束，

	fmt.Printf("c== %v\n", c)
	fmt.Printf("d== %v\n", d)

	//slice 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中int类型的元素是动态的
	doubleArr6 := [2][]int{{1, 34, 5}, {4, 78, 9}}

	doubleArr6[0][0] = 11
	doubleArr6[1][0] = 2

	fmt.Printf("slice声明一个二维数组==    %v\n", doubleArr6)

	/*
		slice有一些简便的操作:
		slice的默认开始位置是0，ar[:n]等价于ar[0:n]
		slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]
		如果从一个数组里面直接获取slice，可以这样ar[:]，因为默认第一个序列是0，第二个是数组的长度，即等价于ar[0:len(ar)]

		slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值


		从概念上面来说slice像一个结构体，这个结构体包含了三个元素：

			一个指针，指向数组中slice指定的开始位置
			长度，即slice的长度
			最大长度，也就是slice开始位置到数组的最后位置的长度

		slice有几个有用的内置函数：
			len 获取slice的长度
			cap 获取slice的最大容量
			append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
			copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数

		append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice
		但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间
		返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；其它引用此数组的slice则不受影响。
	*/

	sliceArrFun := []int{1, 2, 4, 5, 6}

	fmt.Printf("获取slice的长度==    %d\n", len(sliceArrFun))
	fmt.Printf("获取slice的最大容量==    %d\n", cap(sliceArrFun))

	//向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
	sliceArrFun = append(sliceArrFun, 7)

	fmt.Printf(" %v\n", sliceArrFun)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	sliceArrFun2 := make([]int, len(sliceArrFun), (cap(sliceArrFun))*2)

	/* 拷贝 sliceArrFun 的内容到 sliceArrFun2  且返回复制的元素的个数 */
	copy(sliceArrFun2, sliceArrFun)
	fmt.Printf(" %v\n", sliceArrFun2)

}
