/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/16/20, 5:16 PM
 * Author: beihai
 */

package _033_search_in_rotated_sorted_array

// 如果 mid 落在了前一段数值比较大的区间内了，那么 nums[mid] > nums[low]，如果是落在后面一段数值比较小的区间内，nums[mid] ≤ nums[low]
// 如果 mid 落在了后一段数值比较小的区间内了，那么一定有 nums[mid] < nums[high]，如果是落在前面一段数值比较大的区间内，nums[mid] ≤ nums[high]

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
			// 如果中间值在数值大的区间里
		} else if nums[mid] > nums[low] {
			// 如果 target 在 [low，mid) 内
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
			// 如果中间值在数值小的区间里
		} else if nums[mid] < nums[high] {
			// 如果 target 在 [mid，high) 内
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
			// 	特殊情况
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}
