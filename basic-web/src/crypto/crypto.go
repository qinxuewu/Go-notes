/**
 * 存储密码

 目前用的最多的密码存储方案是将明文密码做单向哈希后存储，单向哈希算法有一个特征：无法通过哈希后的摘要(digest)恢复原始数据，这也是“单向”二字的来源。
 常用的单向哈希算法包括SHA-256, SHA-1, MD5等。
 * @author qinxuewu
 * @create 19/5/19上午11:14
 * @since 1.0.0
 */
package main

import (
	"crypto/sha256"
	"io"
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func  main ()  {

	h:=sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Println( h.Sum(nil))

	fmt.Println( "-----------------------------------")


	h2:=sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Println(h2.Sum(nil))


	fmt.Println("-----------------------------------")


	h3:=md5.New()
	io.WriteString(h, "需要加密的密码")
	fmt.Println(h3.Sum(nil))


//	“加盐”的方式来存储密码

	//假设用户名abc，密码123456
	h4:= md5.New()
	io.WriteString(h, "需要加密的密码")

	//pwmd5等于e10adc3949ba59abbe56e057f20f883e
	pwmd5 :=fmt.Sprintf("%x", h4.Sum(nil))

	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h4, salt1)
	io.WriteString(h4, "abc")
	io.WriteString(h4, salt2)
	io.WriteString(h4, pwmd5)

	last :=fmt.Sprintf("%x", h4.Sum(nil))

	fmt.Println("last:  ",last)





}
