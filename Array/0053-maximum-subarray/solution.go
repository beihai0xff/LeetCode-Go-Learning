/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/22/20, 10:09 PM
 * Author: beihai
 */

package _053_maximum_subarray

import . "LeetCode/Util"

/*给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。*/

// 动态规划法
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	currSum, result := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = Max(currSum, 0) + nums[i]
		result = Max(result, currSum)
	}
	return result
}

//
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = Max(res, dp[i])
	}
	return res
}
