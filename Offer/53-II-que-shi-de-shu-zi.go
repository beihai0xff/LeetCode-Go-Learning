package Offer

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
