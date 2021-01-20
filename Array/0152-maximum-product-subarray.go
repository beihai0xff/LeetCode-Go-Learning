/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/20 下午8:47
 * Author: beihai
 */

package Array

import "LeetCode/Util"

// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

// 负数乘以负数，会变成正数，所以解这题的时候我们需要维护两个变量，当前的最大值，以及最小值;
// 最小值可能为负数，但下一步乘以一个负数可能就变成最大值；同时最大值乘以一个负数也可能变成最小值；
// 还要排除 max min为 0 的可能性。
// 最终结果如下：
// 最大值：max = Max(maxT*nums[i], nums[i], minT*nums[i])
// 最小值: min = Min(minT*nums[i], nums[i], maxT*nums[i])
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	max, min, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxT, minT := max, min
		max = Util.Max(maxT*nums[i], Util.Max(nums[i], minT*nums[i]))
		min = Util.Min(minT*nums[i], Util.Min(nums[i], maxT*nums[i]))
		ans = Util.Max(max, ans)
	}
	return ans
}
