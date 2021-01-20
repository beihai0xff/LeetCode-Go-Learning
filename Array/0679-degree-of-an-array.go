/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/20 下午9:21
 * Author: beihai
 */

package Array

import "LeetCode/Util"

// 给定一个非空且只包含非负数的整数数组 nums, 数组的度的定义是指数组里任一元素出现频数的最大值。
//
// 你的任务是找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
//

func findShortestSubArray(nums []int) int {
	left, right, count, res := map[int]int{}, map[int]int{}, map[int]int{}, len(nums)
	// 最短的子数组是将从 x 的第一次出现到最后一次出现的数组
	for i := 0; i < len(nums); i++ {
		if _, ok := left[nums[i]]; !ok {
			left[nums[i]] = i
		}
		right[nums[i]] = i
		count[nums[i]]++
	}
	// 获取最大的度
	var degree int
	for _, val := range count {
		if degree < val {
			degree = val
		}
	}

	// 对于出现次数最多的每个元素 x，right[x] - left[x] + 1 将是我们的候选答案，我们将取这些候选的最小值
	for key, val := range count {
		if degree == val {
			res = Util.Min(res, right[key]-left[key]+1)
		}
	}
	return res
}
