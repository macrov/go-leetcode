package first_missing_positive

func firstMissingPositive(nums []int) int {
	// have 1 ?

	f:=false
	for _, n := range nums {
		if n == 1 {
			f = true
		}
	}
	if !f {
		return 1
	}

	// if nums = [1] return 2
	if len(nums) == 1 {
		return 2
	}

	// transform
	for i:=0;i<len(nums);i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}

	// mark
	for i:=0;i<len(nums);i++ {
		t := abs(nums[i])

		if t == len(nums) {
			nums[0] = -abs(nums[0])
		} else {
			nums[t] = -abs(nums[t])
		}
	}

	// traverse

	for i:=1;i<len(nums);i++ {
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 {
		return len(nums)
	}

	return len(nums)+1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}