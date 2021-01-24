/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/24 下午9:26
 * Author: beihai
 */

package Array

import "sort"

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊n/2⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/

// 排序取中位数
func majorityElement(nums []int) int {
	if len(nums) == 1 || len(nums) == 2 {
		return nums[0]
	}
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 一次遍历，用哈希表存出现的次数
func majorityElement2(nums []int) int {
	if len(nums) == 1 || len(nums) == 2 {
		return nums[0]
	}
	count, max, res := map[int]int{}, 0, 0
	for _, v := range nums {
		count[v]++
		if count[v] > max {
			max = count[v]
			res = v
		}
		if max > len(nums)/2 {
			break
		}
	}
	return res
}
