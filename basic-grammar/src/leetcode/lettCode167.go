package main

import "fmt"

/**
 *
 *  167. 两数之和 II - 输入有序数组\
 *
 * 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
 *
 * 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
 *
 * 返回的下标值（index1 和 index2）不是从零开始的。
 *
 * 你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
 *
 * @author qinxuewu
 * @version 1.00
 * @time  17/4/2019 下午 6:23
 * @email 870439570@qq.com
 */
func main() {
	numbers := []int{2, 7, 11, 15}

	arr := twoSum22(numbers, 9)

	fmt.Println(arr)
}

func twoSum22(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}
	return nil
}
