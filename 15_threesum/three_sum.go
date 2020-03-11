package main

import (
	"fmt"
	"sort"
)

/*
遍历出符合条件的组合很容易，主要难点主要是剪枝。
 */
func threeSum(nums []int) [][]int {
	res := make([][]int,0)
	sort.Ints(nums)

	for l := 0;l < len(nums);l++ {
		// 对于 【-1,-1,0,1,2]  [-1,-1,2]与[-1,0,1]在遍历第一个-1时已经遍历完，第二个-1就不需要再遍历
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		sum := 0 - nums[l]
		m := l + 1
		r := len(nums) -1
		for m < r {
			if nums[m] +  nums[r] == sum {
				res = append(res, []int{nums[l], nums[m], nums[r]})
				// 正好满足条件时，两头去重复
				for m < r && nums[m] == nums[m+1] { m++ }
				for m < r && nums[r] == nums[r-1] { r-- }
				// 去重复后再向中间移动
				m++
				r--
			} else if nums[m] + nums[r] > sum {
				// 后两个数和比需要的大，那r指针向中间移动，先去重
				for m < r && nums[r] == nums[r-1] { r-- }
				// 再移动
				r--
			} else if nums[m] + nums[r] < sum {
				// 后两个数和比需要的小，那m指针向中间移动，先去重
				for m < r && nums[m] == nums[m+1] { m++ }
				// 再移动
				m++
			} else {
				// 剩下的情况就是 l，m两个之和已经大于0，直接跳出循环，l++
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-2,0,0,2,2}))
	fmt.Println(threeSum([]int{-4,-1,-1,0,1,2,6}))
	fmt.Println(threeSum([]int{0,0,0,0}))
	fmt.Println(threeSum([]int{1,-1,-1,0}))
	fmt.Println(threeSum([]int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}))
}
