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

// 给定排序好的数组如[1,1,1,2,2,4,4,6,7] 对其进行重复数字的过滤，过滤出来的数字放在数组的末端，
// 如对上面数组过滤后的结果为[1,2,4,6,7,(1,1,2,4)]
// 括号当中的数字即为过滤出来的数字，放在数组最末端，
// 这里括号仅为了方便阅读，括号里的数字不要求排序，要求空间复杂度为O（1）
// [1,1,1,2,2,4,4,6,7]
// [1,2,1,1,2,4,4,6,7]
// [1,2,4,1,2,1,4,6,7]
func removeDuplicates2(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	left, right, port := 1, 1, nums[0]
	for right < len(nums) {
		if nums[right] > port {
			port = nums[right]
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return
}
