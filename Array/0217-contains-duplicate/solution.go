/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午7:09
 * Author: beihai
 */

package _217_contains_duplicate

// 哈希表存储遍历过的值
func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, value := range nums {
		if record[value] {
			return true
		}
		record[value] = true
	}
	return false
}
