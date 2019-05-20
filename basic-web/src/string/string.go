/**
 * Go标准库中的strings和strconv两个包中的函数 处理字符串
 * @author qinxuewu
 * @create 19/5/14下午10:41
 * @since 1.0.0
 */
package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {

	//判断字字符串s中是否包含substr，返回bool值
	fmt.Println(strings.Contains("qxw","q")) // truee
	fmt.Println(strings.Contains("qxw","aaa")) //false
	fmt.Println(strings.Contains("qxw","")) //true
	fmt.Println(strings.Contains("","")) //true

//	字符串链接，
   s:=[]string{"a","b","c"}
   fmt.Println(strings.Join(s,","))

//   在字符串s中查找sep所在的位置，返回位置值，找不到返回-1

   fmt.Println(strings.Index("qinxuewu","xue"))
   fmt.Println(strings.Index("qinxuewu","a"))

//   重复s字符串count次，最后返回重复的字符串
   fmt.Println("ba"+strings.Repeat("na",2))

//   在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
   fmt.Println(strings.Replace("qxw xue wu","q","Q",2))
   fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

//   把s字符串按照sep分割，返回slice
   fmt.Printf("%q\n",strings.Split("a,b,c",","))
   fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
   fmt.Printf("%q\n", strings.Split(" xyz ", ""))
   fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	//Output:["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]


//在s字符串的头部和尾部去除cutset指定的字符串

  fmt.Printf("[%q]",strings.Trim("!!! qxw !!!","! "))

//   去除s字符串的空格符，并且按照空格分割返回slice

 fmt.Printf("Fields are:%q\n",strings.Fields(" foo bar baz "))


	//字符串转换
	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str:=make([]byte,0,100)
	str=strconv.AppendInt(str,4567,10)
	str=strconv.AppendBool(str,false)
	str=strconv.AppendQuote(str,"abcdefg")
	str=strconv.AppendQuoteRune(str,'单')
	fmt.Println("字符串转换",string(str))

//	Format 系列函数把其他类型的转换为字符串

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)  // false 123.23 1234 12345 1023

//	Parse 系列函数把字符串转换为其他类型

	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)   //false 123.23 1234 12345 1023

}


func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}