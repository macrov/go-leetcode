package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int, target int) [][]int {
	res := make([][]int,0)
	for l := 0;l < len(nums);l++ {
		// 对于 【-1,-1,0,1,2]  [-1,-1,2]与[-1,0,1]在遍历第一个-1时已经遍历完，第二个-1就不需要再遍历
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		sum := target - nums[l]
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

func fourSum(nums []int, target int) [][]int {
	res := make([][]int,0)
	sort.Ints(nums)
	for i:=0;i<len(nums)-3;i++ {
		if i>0 && nums[i] == nums[i-1] {
			continue
		}
		t := threeSum(nums[i+1:], target - nums[i])
		for j:=0;j<len(t);j++ {
			res = append(res, append([]int{nums[i]}, t[j]...))
		}
	}
	return res
}

func main() {
	fmt.Println(fourSum([]int{1,0,-1,0,-2,2}, 0))
}
