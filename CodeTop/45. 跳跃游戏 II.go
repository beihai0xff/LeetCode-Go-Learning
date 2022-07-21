package CodeTop

func jump(nums []int) int {
	maxPos, index, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == index {
			index = maxPos
			steps++
		}
	}

	return steps
}
