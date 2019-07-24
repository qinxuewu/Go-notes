package main

import "fmt"

/**
 *  两个数组的交集  给定两个数组，编写一个函数来计算它们的交集。
 *
 *  输入: nums1 = [1,2,2,1], nums2 = [2,2]
 *  输出: [2]
 *
 *  输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 *  输出: [9,4]
 *
 *  输出结果中的每个元素一定是唯一的。
 *  我们可以不考虑输出结果的顺序。
 *
 * @author qxw
 * @version 1.00
 * @time  24/7/2019 上午 11:40
 */

func main()  {
	nums1:=[]int{1,2,2,1}
	nums2:=[]int{2,2}
	data:=intersection(nums1,nums2)
	fmt.Println(data)
}

// 使用map的key唯一的特性 去重
func intersection(nums1 []int, nums2 []int) []int {
	set:=make(map[int]int)
	res:=make([]int,0)  // 结果集

	for _,v:=range  nums1 {
		// 数组中存储到map的k
		set[v]=1
	}

	for _, v := range nums2 {
		// 如果set名称的map中存在当当前key返回ture 并且key对应的value为1
		if _, ok := set[v]; ok && set[v] == 1{
			res = append(res, v)
			set[v] = 0
		}
	}
	return  res
}