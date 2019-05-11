/**
 * 模板处理
 * @author qinxuewu
 * @create 19/5/11上午12:00
 * @since 1.0.0
 */
package main

import (
	"html/template"
	"os"
)

func main()  {
	//test()
	test2()
}

type Preson struct {
	UserName string
}

func test()  {
	t:=template.New("fieldname example")

	//Go语言的模板通过{{}}来包含需要在渲染时被替换的字段 {{.}}表示当前的对象
	t,_=t.Parse("hello  {{.UserName}}!")
	p:=Preson{UserName:"Astaxie"}
	t.Execute(os.Stdout,p)
}

/**
	输出嵌套字段内容
	如果字段里面还有对象 使用{{with …}}…{{end}}和{{range …}}{{end}}来进行数据的输出
	{{range}} 这个和Go语法里面的range类似，循环操作数据
	{{with}}操作是指当前对象的值，类似上下文的概念
 */
type Friend struct {
	Fname string
}

type Person2 struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func test2()  {
	f1:=Friend{Fname:"minux.ma"}
	f2:=Friend{Fname:"qddddd"}
	t:=template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
			{{range .Emails}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)

	p := Person2{
				UserName: "Astaxie",
				Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
				Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}