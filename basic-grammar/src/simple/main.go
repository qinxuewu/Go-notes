package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 聊天机器人,启动之后不断询问

func main()  {
	// 从控制台读取数据
	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字:")
	// 读取到\n 就停止
	input,err:=inputReader.ReadString('\n')
	if err !=nil{
		fmt.Printf("系统异常: %s\n",err)
		// 异常退出
		os.Exit(1)
	}else {
		// 切皮操作删除最后的 \n
		name:=input[:len(input)-1]
		fmt.Printf("你好 ,%s! 我有说吗可以帮助你的吗?\n",name)
	}

	for  {
		input,err=inputReader.ReadString('\n')
		if err !=nil {
			fmt.Printf("系统异常:%s\n",err)
			continue
		}
		input=input[:len(input)-1]
		// 全部转换小写
		input=strings.ToLower(input)
		switch  input {
		case "":
			continue
		case "nothing","bye":
			fmt.Println("输入特定的词语 结束程序!")
		    os.Exit(0)
		default:
			fmt.Println("Sorry,我不理解你的意思.")

		}
	}
}
