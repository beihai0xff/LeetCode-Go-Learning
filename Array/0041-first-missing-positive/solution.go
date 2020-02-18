/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/18/20, 9:38 PM
 * Author: beihai
 */

package _041_first_missing_positive

import (
	. "LeetCode/Util"
)

/*
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
*/

// 官解，将自己做为 bitmap：https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode/
func firstMissingPositive(nums []int) int {

	lens, contains := len(nums), 0
	for i := 0; i < lens; i++ {
		if nums[i] == 1 {
			contains++
		}
	}
	if contains == 0 {
		return 1
	}

	for i := 0; i < lens; i++ {
		if nums[i] <= 0 || nums[i] > lens {
			nums[i] = 1
		}
	}

	for i := 0; i < lens; i++ {
		// 避免重复元素
		a := AbsInt(nums[i])
		if a == lens {
			nums[0] = -AbsInt(nums[0])
		} else if a < lens {
			nums[a] = -AbsInt(nums[a])
		} else {
			continue
		}
	}

	for i := 1; i < lens; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return lens
	}
	return lens + 1
}

// 把数据放在一个 map 里面查
func firstMissingPositive2(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v] = v
	}
	for index := 1; index < len(nums)+1; index++ {
		if _, ok := numMap[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}
