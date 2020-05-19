package combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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
		if i>0 && candidates[i-1] == candidates[i] {
			continue
		}
		c := candidates[i]
		if c > target {
			return
		}
		if c <= target {
			bt(candidates[i+1:], target - c, append(nums, c), res)
		}
	}
}