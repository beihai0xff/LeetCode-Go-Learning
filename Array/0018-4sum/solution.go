/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 1:40 PM
 * Author: beihai
 */

package _018_4sum

import "sort"

// 暴力查找，做了优化，速度很快
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	lens := len(nums)
	// 定位第一个数字 nums[i]
	for i := 0; i < lens-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		min := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if min > target {
			break
		}
		max := nums[i] + nums[lens-1] + nums[lens-2] + nums[lens-3]
		if max < target {
			continue
		}
		// 定位第二个数字 nums[j]
		for j := i + 1; j < lens-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			m := j + 1
			n := lens - 1
			min2 := nums[j] + nums[m] + nums[i+2] + nums[i]
			if min2 > target {
				continue
			}
			max2 := nums[i] + nums[j] + nums[n] + nums[lens-2]
			if max2 < target {
				continue
			}
			// 查找 nums[m],nums[n]
			for m < n {
				temp := nums[i] + nums[j] + nums[m] + nums[n]
				if temp == target {
					result = append(result, []int{nums[i], nums[j], nums[m], nums[n]})
					// 去重
					m++
					for m < n && nums[m] == nums[m-1] {
						m++
					}
					n--
					for m < n && nums[n] == nums[n+1] {
						n--
					}
				} else if temp > target {
					n--
				} else {
					m++
				}
			}
		}
	}
	return result
}

// 和 15 题的思路差不多，提前计算好任意 3 个数字之和，用 map 保存
func fourSum1(nums []int, target int) [][]int {
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
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*3+uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*3+uniqNums[i] == target) && counter[uniqNums[j]] > 2 {
				result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i]*2 == target) && counter[uniqNums[j]] > 1 && counter[uniqNums[i]] > 1 {
				result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[j]*2+uniqNums[i]+uniqNums[k] == target) && counter[uniqNums[j]] > 1 {
					result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[k]*2+uniqNums[i]+uniqNums[j] == target) && counter[uniqNums[k]] > 1 {
					result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0 {
					result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return result
}
