/**
 * 正则处理
 * @author qinxuewu
 * @create 19/5/10下午11:29
 * @since 1.0.0
 */
package main

import (
	"regexp"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func main()  {
	//fmt.Println(IsIP("192.168.1.11"))
	//fmt.Println(IsString("双方都"))
	//test()
	test2()
	
}

//验证一个输入是不是IP地址
func  IsIP(ip string)(b bool) {
	if m,_:=regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip);!m{
		return false

	}
	return  true
}

// 判断一个字符串是不是 数字
func IsString(s string)(b bool) {
	 if m,_:=regexp.MatchString("^[0-9]",s);m{
	 	return  true
	 }else {
	 	return  false
	 }
}


/**
	通过正则获取内容
	首先是Compile，它会解析正则表达式是否合法
	正确，那么就会返回一个Regexp，然后就可以利用返回的Regexp在任意的字符串上面执行需要的操作
 */
func  test()  {
	resp,err:=http.Get("http://www.baidu.com")
	if err !=nil{
		fmt.Println(" http get error......")
	}
	defer  resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)

	if err !=nil{
		fmt.Println("http red error...")
		return
	}

	src:=string(body)

	// 将HTML标签全转换成小写
	re,_:=regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除sytle
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")


	//去除Sscript
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))

}

func  test2()  {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置,开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的，第三个是第二个()里面的
	//下面的输出第一个元素是"am learning Go language"
	//第二个元素是" learning Go "，注意包含空格的输出
	//第三个元素是"uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}