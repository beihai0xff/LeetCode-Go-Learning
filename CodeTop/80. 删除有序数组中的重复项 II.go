package CodeTop

func removeDuplicates80(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
