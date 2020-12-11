/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午1:13
 * Author: beihai
 */

package _977_squares_of_a_sorted_array

import "LeetCode/Util"

// 双指针：用 2 个指针分别指向原数组的首尾，分别计算平方值，然后比较两者大小，大的放在最终数组的后面。
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	// 需要包含 left == right 的情况
	for offset := len(res) - 1; left <= right; offset-- {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[offset] = nums[left] * nums[left]
			left++
		} else {
			res[offset] = nums[right] * nums[right]
			right--
		}
	}
	return res
}

// 相关题目：有序数组每个数平方后，不同数字的个数

// 平方之后在两边的数据较大，中间的数据较小
// 双指针，从两侧向中间扫描，并去重
func squareUniqueNum(nums []int) int {
	left, right, res, deleted := 0, len(nums)-1, 0, 0
	for left != right {
		if Util.AbsInt(nums[left]) > Util.AbsInt(nums[right]) {
			// 去重
			if deleted != Util.AbsInt(nums[left]) {
				deleted = Util.AbsInt(nums[left])
				res++
			}
			left++
		} else if Util.AbsInt(nums[left]) == Util.AbsInt(nums[right]) {
			if deleted != Util.AbsInt(nums[left]) {
				deleted = Util.AbsInt(nums[left])
				res++
			}
			left++
			right--
		} else {
			if deleted != Util.AbsInt(nums[right]) {
				deleted = Util.AbsInt(nums[right])
				res++
			}
			right--
		}
	}
	return res
}
