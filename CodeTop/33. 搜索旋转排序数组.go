package CodeTop

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 左侧有序
		if nums[left] <= nums[mid] {
			// 判断 target 在哪个区间
			if nums[left] <= target && target < nums[mid] { // 左区间
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧有序
			if nums[mid] < target && target <= nums[right] { // 右区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
