> [goquery官方文档地址](https://godoc.org/github.com/PuerkitoBio/goquery)
> 会用jquery的，goquery基本可以1分钟上手

### goquery常用方法
```
Eq(index int) *Selection 	//根据索引获取某个节点集
First() *Selection 			//获取第一个子节点集
Last() *Selection	 		//获取最后一个子节点集
Next() *Selection 			//获取下一个兄弟节点集
NextAll() *Selection 		//获取后面所有兄弟节点集
Prev() *Selection			//前一个兄弟节点集
Get(index int) *html.Node	//根据索引获取一个节点
Index() int 				//返回选择对象中第一个元素的位置
Slice(start, end int) *Selection //根据起始位置获取子节点集
Each(f func(int, *Selection)) *Selection //遍历
EachWithBreak(f func(int, *Selection) bool) *Selection //可中断遍历
Map(f func(int, *Selection) string) (result []string) //返回字符串数组
Attr(), RemoveAttr(), SetAttr() //获取，移除，设置属性的值
AddClass(), HasClass(), RemoveClass(), ToggleClass()
Html() //获取该节点的html
Length() //返回该Selection的元素个数
Text() //获取该节点的文本值
Children() //返回selection中各个节点下的孩子节点
Contents() //获取当前节点下的所有节点
Find() //查找获取当前匹配的元素
Next() //下一个元素
Prev() //上一个元素
```

##  参考
- https://blog.csdn.net/jeffrey11223/article/details/79318856
- 采集地址： http://zuikzy.cc/?m=vod-type-id-1.html

