/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午3:18
 * Author: beihai
 */

package Util

// 二分搜索，返回查找元素所在的下标
func binarySearch(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1

		if k > nums[mid] {
			left = mid + 1
		} else if k < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
