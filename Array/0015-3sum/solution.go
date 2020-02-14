/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/13/20, 1:54 PM
 * Author: beihai
 */

package leetcode

import "sort"

// 双指针 + 求双数和，提前计算好任意 3 个数字之和，用 map 保存
func threeSum(nums []int) [][]int {
	var result [][]int
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		// 全为 0
		if (uniqNums[i] == 0) && counter[uniqNums[i]] >= 3 {
			result = append(result, []int{0, 0, 0})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			// 有两个 uniqNums[i]
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			// 有两个 uniqNums[j]
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			// 三个全不同
			// 计算第三个数
			c := 0 - uniqNums[i] - uniqNums[j]
			// 查找
			if c > uniqNums[j] && counter[c] > 0 {
				result = append(result, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return result
}

// 排序 + 双指针
func threeSum1(nums []int) [][]int {
	lens := len(nums)
	var res [][]int
	// 排序
	sort.Ints(nums)
	for i := 0; i < lens; i++ {
		// 有序数组，大于 0 之后就不会有了
		if nums[i] > 0 {
			return res
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high := i+1, lens-1
		// 双指针
		for low < high {
			// 去重
			if low > i+1 && nums[low] == nums[low-1] {
				low++
				continue
			}
			if high < lens-1 && nums[high] == nums[high+1] {
				high--
				continue
			}
			result := nums[i] + nums[low] + nums[high]
			if result == 0 {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				low++
				high--
			} else if result < 0 {
				low++
			} else {
				high--
			}
		}
	}
	return res
}
