package main

import "fmt"

// 两数之和

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

**/
func main()  {
	nums := []int{2, 7, 11, 15}

	res:=twoSum2(nums,17)
	fmt.Println(res)
}


func twoSum(nums []int, target int) []int {
	res:=make([]int,2)
	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums) ; j++ {
			if nums[i]+nums[j] == target {
				res[0]=i;
				res[1]=j;
				return res
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	rut:=make([]int,2)
	m := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		res:=target-nums[i]
		value, ok := m[res]
		if ok {
			rut[0]=value
			rut[1]=i
			return rut
		}
		m[nums[i]]=i
	}
	return nil
}