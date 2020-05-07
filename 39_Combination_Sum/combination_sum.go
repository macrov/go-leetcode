package combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int,0, 0)
	bt(candidates, target, []int{}, &res)
	return res
}

func bt(candidates []int, target int, nums []int, res *[][]int) {
	if target == 0 {
		cb := make([]int,len(nums))
		copy(cb, nums)
		*res = append(*res, cb)
		return
	}

	for i:=0;i<len(candidates);i++{
		c := candidates[i]
		if c > target {
			return
		}
		if c <= target {
			bt(candidates[i:], target - c, append(nums, c), res)
		}
	}
}