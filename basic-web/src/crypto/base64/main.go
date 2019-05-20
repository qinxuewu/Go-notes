/**
 * base64加解密
 * @author qinxuewu
 * @create 19/5/20下午8:48
 * @since 1.0.0
 */
package main

import (
	"encoding/base64"
	"fmt"
)

// 加密
func base64Encode( src []byte)[] byte  {
	return  [] byte(base64.StdEncoding.EncodeToString(src))
}

// 解密

func base64Decode(src []byte) ([]byte,error)  {
	return  base64.StdEncoding.DecodeString(string(src))
}

func main()  {
  hello:="努力！ 奋斗 换工作"
  debyte:=base64Encode([]byte(hello))
  fmt.Println(" 加密后：",debyte)

  enbyte,err:=base64Decode(debyte)
  if err!=nil{
  	fmt.Println(err.Error())
  }

  if hello !=string(enbyte){
  	fmt.Println("解密的字符串和原始字符不匹配")
  }

  fmt.Println(string(enbyte))
}


