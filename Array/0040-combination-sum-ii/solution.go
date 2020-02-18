/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/18/20, 9:07 PM
 * Author: beihai
 */

package _040_combination_sum_ii

import "sort"

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：
	所有数字（包括目标数）都是正整数。
	解集不能包含重复的组合。
*/

// 与 39 题一致，但是每个数字在每个组合中只能使用一次
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	sort.Ints(candidates)
	dfs(candidates, target, 0, c, &res)
	return res
}

func dfs(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // 去重，本次相同的值在上一次的 dfs 中已经取到
			continue
		}
		c = append(c, nums[i])
		dfs(nums, target-nums[i], i+1, c, res)
		c = c[:len(c)-1]
	}
}
