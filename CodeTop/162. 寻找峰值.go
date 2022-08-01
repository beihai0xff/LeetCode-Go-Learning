package CodeTop

func findPeakElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	var res int
	for i, v := range nums {
		if v > nums[res] {
			res = i
		}
	}

	return res
}
