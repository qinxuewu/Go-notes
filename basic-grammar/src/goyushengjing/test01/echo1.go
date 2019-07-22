/**
 *
 * @author qinxuewu
 * @create 19/7/1下午9:36
 * @since 1.0.0
 */
package main

import (
	"os"
	"fmt"
)

func main()  {
	
	var s,sep string
	for i:=1;i<len(os.Args);i++ {
		s += sep + os.Args[i]
		sep =" "
	}
	fmt.Println(s)
}