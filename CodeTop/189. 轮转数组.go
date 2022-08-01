package CodeTop

func rotate189(nums []int, k int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	k %= len(nums)
	reverseRange(nums, 0, len(nums)-1)
	reverseRange(nums, 0, k)
	reverseRange(nums, k+1, len(nums)-1)
}

func reverseRange(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
