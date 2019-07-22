/**
 *
 * @author qinxuewu
 * @create 19/7/3下午10:30
 * @since 1.0.0
 */
package main

import (
	"Go-notes/basic-grammar/src/queue"
	"fmt"
)

func main()  {

	q:=queue.Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
