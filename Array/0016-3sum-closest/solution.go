/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/13/20, 2:57 PM
 * Author: beihai
 */

package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	lens, res, diff := len(nums), 0, math.MaxInt32
	if lens > 2 {
		// 排序
		sort.Ints(nums)
		for i := 0; i < lens-2; i++ {
			// 双指针求和
			for low, high := i+1, lens-1; low < high; {
				sum := nums[i] + nums[low] + nums[high]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					high--
				} else {
					low++
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
