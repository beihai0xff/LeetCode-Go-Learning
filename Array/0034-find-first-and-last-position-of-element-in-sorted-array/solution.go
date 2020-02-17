/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/16/20, 7:45 PM
 * Author: beihai
 */

package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"sort"
)

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	leftIndex := extremeInsertionIndex(nums, target, true)
	if leftIndex == len(nums) || nums[leftIndex] != target {
		return result
	}
	result[0] = leftIndex
	result[1] = extremeInsertionIndex(nums, target, false) - 1
	return result
}

func extremeInsertionIndex(nums []int, target int, left bool) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > target || (left && nums[mid] == target) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 使用标准库的二分查找法，最差情况下时间复杂度为 O(n)
func searchRange1(nums []int, target int) []int {

	i := sort.SearchInts(nums, target) // Sort Search
	if len(nums) == 0 || i == len(nums) || nums[i] != target {
		return []int{-1, -1}
	}

	var j int // 找右边界
	for j = i; j < len(nums) && nums[i] == nums[j]; j++ {
	}
	if i < len(nums) || nums[i] == target {
		return []int{i, j - 1}
	}
	return []int{-1, -1}

}
