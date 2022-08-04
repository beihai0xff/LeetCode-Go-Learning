package CodeTop

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left, right := 1, 1
	for right < len(nums) {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
		right++
	}

	return left
}
