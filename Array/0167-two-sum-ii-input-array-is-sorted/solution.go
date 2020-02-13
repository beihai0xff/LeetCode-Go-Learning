/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/13/20, 12:16 PM
 * Author: beihai
 */

package leetcode

/*使用双指针，一个指针指向值较小的元素，一个指针指向值较大的元素。指向较小元素的指针从头向尾遍历，指向较大元素的指针从尾向头遍历。

如果两个指针指向元素的和 sum == targetsum==target，那么得到要求的结果；
如果 sum > targetsum>target，移动较大的元素，使 sumsum 变小一些；
如果 sum < targetsum<target，移动较小的元素，使 sumsum 变大一些。

数组中的元素最多遍历一次，时间复杂度为 O(N)。只使用了两个额外变量，空间复杂度为 O(1)。*/

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return nil
}

// 题 0001 的解法
func twoSum1(numbers []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		another := target - numbers[i]
		if _, ok := m[another]; ok {
			return []int{m[another] + 1, i + 1}
		}
		m[numbers[i]] = i
	}
	return nil
}
