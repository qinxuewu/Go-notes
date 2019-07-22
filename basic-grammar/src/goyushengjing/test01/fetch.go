/**
 *
 * @author qinxuewu
 * @create 19/7/3下午9:56
 * @since 1.0.0
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 读取控制台传入的url
	for _, url := range os.Args[1:] {
		// 发送get请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// 读取失败 需要退出
			os.Exit(1)
		}
		// 从response中读取到全部内容
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}