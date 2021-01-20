/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/20 下午8:18
 * Author: beihai
 */

package Array

import (
	"LeetCode/Util"
	"math"
	"sort"
)

// 给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

// 排序求解
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	// 分别求出三个最大正数的乘积，以及两个最小负数与最大正数的乘积，二者之间的最大值即为所求答案
	return Util.Max(nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3], nums[0]*nums[1]*nums[len(nums)-1])
}

// 线性扫描求出上面五个值
func maximumProduct2(nums []int) int {
	// 最小的和第二小的
	min1, min2 := math.MaxInt64, math.MaxInt64
	// 最大的、第二大的和第三大的
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums { // v < min1 < min2
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 { // min1 < v < min2
			min2 = v
		}
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	return Util.Max(min1*min2*max1, max1*max2*max3)
}
