package Offer

import "sort"

func search(nums []int, target int) int {
	// 在数组中寻找第一个大于等于 target 的下标
	leftIndex := sort.SearchInts(nums, target)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return 0
	}

	// 在数组中寻找第一个大于 target 的下标
	rightIndex := sort.SearchInts(nums, target+1)
	return rightIndex - leftIndex
}
