/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/25 下午10:06
 * Author: beihai
 */

package Array

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
*/

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))

	start := 0
	end := 0
	pi := 0

	for _, v := range nums {
		end = pi
		start = 0
		for start < end {
			mid := start + (end-start)/2
			if top[mid] >= v {
				end = mid
			} else {
				start = mid + 1
			}
		}

		if start == pi {
			pi++
		}

		top[start] = v

	}

	return pi
}
