package CodeTop

func singleNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	var res int
	for _, num := range nums {
		res ^= num
	}

	return res
}
