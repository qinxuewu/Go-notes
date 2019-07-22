/**
 *  实现先进先出的队列
 * @author qinxuewu
 * @create 19/7/3下午10:27
 * @since 1.0.0
 */
package queue

type Queue []int

// q指向的sicle改变
func (q  *Queue) Push(v int) {
	*q =append(*q,v)
}

func (q *Queue) Pop() int  {
	head :=(*q)[0]
	*q=(*q)[1:]
	return  head
}
func (q *Queue) IsEmpty() bool  {

	return len(*q) == 0
}
