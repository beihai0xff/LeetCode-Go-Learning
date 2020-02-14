/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 4:04 PM
 * Author: beihai
 */

package _026_remove_duplicates_from_sorted_array

// 双指针，将删除的元素移动到数组后面的空间
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[start] {
			start++
			if i-start > 0 {
				nums[start] = nums[i]
			}
		}
	}
	return start + 1
}
