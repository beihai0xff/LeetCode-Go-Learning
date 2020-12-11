/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午3:18
 * Author: beihai
 */

package Util

// 二分搜索，返回查找元素所在的下标

// 三点注意事项：
// 1. 循环退出条件，注意是 left <= right，而不是 left < right。
// 2. mid 的取值，mid := left + (right-left)»1
// 3. left 和 right 的更新: left = mid + 1，right = mid - 1。
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1

		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 二分搜索四个变种：
// 1. 查找第一个与 target 相等的元素，时间复杂度 O(logn)
// 2. 查找最后一个与 target 相等的元素，时间复杂度 O(logn)
// 3. 查找第一个大于等于 target 的元素，时间复杂度 O(logn)
// 4. 查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
