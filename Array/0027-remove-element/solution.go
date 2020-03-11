/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 4:39 PM
 * Author: beihai
 */

package _027_remove_element

// //  26、27、80、283 原理相通，双指针法
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != start {
				nums[i], nums[start] = nums[start], nums[i]
			}
			start++
		}
	}
	return start
}
