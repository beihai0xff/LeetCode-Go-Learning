package Offer

func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	head, tail := 0, len(nums)-1
	for head < tail && head < len(nums) && tail >= 0 {
		for head < len(nums) && nums[head]%2 == 1 {
			head++
		}
		for tail >= 0 && nums[tail]%2 == 0 {
			tail--
		}
		if head < tail && head < len(nums) && tail >= 0 {
			nums[head], nums[tail] = nums[tail], nums[head]
			head++
			tail--
		}
	}
	return nums
}
