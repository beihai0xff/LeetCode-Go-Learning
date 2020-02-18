/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/18/20, 3:59 PM
 * Author: beihai
 */

package _039_combination_sum

import (
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：
	所有数字（包括 target）都是正整数。
	解集不能包含重复的组合。
*/

// 回溯算法
func combinationSum(candidates []int, target int) [][]int {
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
			// append 为弱复制，用临时变量更稳定
			b := make([]int, len(c))
			// copy 复制为值复制，= 复制为指针复制。
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		c = append(c, nums[i])
		dfs(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}

// 递归穷举
func combinationSum2(candidates []int, target int) [][]int {
	var comb [][]int
	for i := range candidates {
		if candidates[i] == target {
			comb = append(comb, []int{candidates[i]})
		} else if candidates[i] < target {
			rtn := combinationSum(candidates[i:], target-candidates[i])
			for j := range rtn {
				if len(rtn[j]) == 0 {
					continue
				}
				comb = append(comb, append([]int{candidates[i]}, rtn[j]...))
			}
		}
	}
	return comb
}
