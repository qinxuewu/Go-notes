package main

import "fmt"

// go运算符操作

/**
Go 语言内置的运算符有：

		算术运算符
		关系运算符
		逻辑运算符
		位运算符
		赋值运算符
		其他运算符
*/

func main() {
	//test1()
	//test2()

}

//算术运算符
func test1() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)

	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)

	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

}

//关系运算符
func test2() {
	var a int = 21
	var b int = 10

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}
	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}
}

//逻辑运算符
func test3() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

/**
位运算符

&   按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。  (A & B) 结果为 12, 二进制为 0000 1100
|   按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
^   按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
<<  左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。
>>   右移运算符">>"是双目运算符。右移n位就是除以2的n次方
*/
func test4() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}

//赋值运算符

/**
=	简单的赋值运算符，将一个表达式的值赋给一个左值
+=	相加后再赋值
-=	相减后再赋值
*=	相乘后再赋值
/=	相除后再赋值
%=	求余后再赋值
<<=	左移后赋值
>>=	右移后赋值
&=	按位与后赋值
^=	按位异或后赋值
|=	按位或后赋值
*/
func test5() {
	var a int = 21
	var c int

	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c) //21

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c) //42

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c) //21

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c) // 21*21=441

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c) //   441/21=21

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c) //800 	C=200<<2

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c) //200  C=800>>2

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c) //0  C=200 & 2

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c) //2 	C=0 ^ 2

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c) //2 C=2 | 2
}

/*
	其他运算符

	&	返回变量存储地址   &a; 将给出变量的实际地址。
	*	指针变量。		*a; 是一个指针变量
*/
func test6() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a) // int
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b) // int32
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c) // float32

	/*  & 和 * 运算符实例 */
	ptr = &a                        /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)    // 4
	fmt.Printf("*ptr 为 %d\n", *ptr) //4
}
