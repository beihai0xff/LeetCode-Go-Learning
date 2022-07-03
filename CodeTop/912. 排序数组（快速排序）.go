package CodeTop

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	// 递归终止条件
	if start > end {
		return
	}
	// 假设 pivot 为最左侧元素
	left, right, pivot := start, end, nums[start]
	for left < right {
		// 先从右侧开始
		for left < right && nums[right] >= pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	// 交换 pivot 为分界线位置
	nums[left], nums[start] = nums[start], nums[left]
	// pivot 元素本身不用递归
	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}
