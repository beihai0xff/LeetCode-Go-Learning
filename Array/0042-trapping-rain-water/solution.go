/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/19/20, 12:30 PM
 * Author: beihai
 */

package _042_trapping_rain_water

import . "LeetCode/Util"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

// 下面四种方法源自官方题解：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/

// 暴力解法：对数组中的每个元素，找出下雨后水能达到的最高位置，等于两边最大高度的较小值减去当前高度的值。
func trap(height []int) int {
	result, lens := 0, len(height)
	for i := 1; i < lens-1; i++ {
		leftMax, rightMax := 0, 0
		for j := i; j >= 0; j-- {
			leftMax = Max(height[j], leftMax)
		}
		for j := i; j < lens; j++ {
			rightMax = Max(height[j], rightMax)
		}
		result += Min(leftMax, rightMax) - height[i]
	}
	return result
}

// 动态规划
func trap1(height []int) int {
	result, lens := 0, len(height)
	// 为空
	if lens == 0 {
		return 0
	}
	leftMax, rightMax := make([]int, lens), make([]int, lens)
	leftMax[0] = height[0]
	for i := 1; i < lens; i++ {
		leftMax[i] = Max(height[i], leftMax[i-1])
	}
	rightMax[lens-1] = height[lens-1]
	for i := lens - 2; i >= 0; i-- {
		rightMax[i] = Max(height[i], rightMax[i+1])
	}
	for i := 1; i < lens; i++ {
		result += Min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}

// 双指针
func trap2(height []int) int {
	result, left, right, leftMax, rightMax := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > leftMax {
				// 设置左边最高柱子
				leftMax = height[left]
			} else {
				// 右边必定有柱子挡水，所以，遇到所有值小于等于leftMax的，全部加入水池
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				// 设置右边最高柱子
				rightMax = height[right]
			} else {
				// 左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}
