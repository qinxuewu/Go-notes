/**
 *
 *  文件操作
 * @author qinxuewu
 * @create 19/5/13下午8:35
 * @since 1.0.0
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//fileTest()
	//test2()
	//test3()
	listPath("F:\\ideWork\\githubWork\\Go-notes\\basic-web\\src")
}

//目录操作
func fileTest() {
	//创建名称为name的目录，权限设置是perm，例如0777
	os.Mkdir("/Users/qinxuewu/Desktop/GiWork/Go-notes/astaxie", 0777)
	//根据path创建多级子目录
	os.MkdirAll("/Users/qinxuewu/Desktop/GiWork/Go-notes/astaxie/test1/test2", 0777)
	//删除名称为name的目录  当目录下有文件或者其他目录时会出错
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
	}
	//根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。
	os.RemoveAll("/Users/qinxuewu/Desktop/GiWork/Go-notes/astaxie")

}

/**
	文件操作
	Create: 根据提供的文件名创建新的文件，返回一个文件对象，
    NewFile:  根据文件描述符创建相应的文件，返回一个文件对象
    Open: 该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile。
    OpenFile: 打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限
    Write:  写入byte类型的信息到文件
    WriteAt:  在指定位置开始写入byte类型的信息
    WriteString:  写入string信息到文件

*/

// 写文件
func test2() {
	userFile := "astaie.text"
	//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的
	foit, err := os.Create(userFile)
	if err != nil {
		fmt.Printf(userFile, err)
		return
	}
	defer foit.Close()

	for i := 0; i < 10; i++ {
		//写入string信息到文件
		foit.WriteString("Just a test!\r\n")
		//写入byte类型的信息到文件
		foit.Write([]byte("Just a test!\r\n"))
	}
}

// 读文件

func test3() {
	userFile := "astaie.text"
	f1, err := os.Open(userFile)
	if err != nil {
		fmt.Printf(userFile, err)
	}
	defer f1.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := f1.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

// 遍历指定目录下 的文件和目录
func listPath(path string) {
	list, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range list {
		fmt.Println("名称，类型,路径", f.Name(), f.IsDir(), path+f.Name())
	}
}

// 读取指定文件路径的内容
func ReadFileText(path string) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(b)
	fmt.Println(text)
}

// 写入指定文件的内容
func WriteFileText(path, text string) {
	b := []byte(text)

	// WriteFile 向文件 path 中写入数据 data
	// 如果文件不存在，则以 perm 权限创建该文件
	// 如果文件存在，则先清空文件，然后再写入
	err := ioutil.WriteFile(path, b, 0666)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}

// 编写程序输出某目录下的所有文件（包括子目录）
func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"\\"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
