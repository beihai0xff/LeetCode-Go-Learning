package CodeTop

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	// 寻找第一个降序前的数字，为 i-1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			// 寻找第一个比nums[i-1]大的数字
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					// 翻转 i 到最后的数字
					reverse(nums[i:])
					return
				}
			}
		}
	}

	reverse(nums)
	return
}
