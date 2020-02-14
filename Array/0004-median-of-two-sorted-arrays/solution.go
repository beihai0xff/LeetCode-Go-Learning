/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/13/20, 10:56 PM
 * Author: beihai
 */

package leetcode

import (
	. "LeetCode/Util"
	"sort"
)

// 二分法，使得数组左侧小于右侧
// 题解思路参考：https://github.com/halfrost/LeetCode-Go/tree/master/Algorithms/0004.%20Median%20of%20Two%20Sorted%20Arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，输出最终结果
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = Max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	// 如果为奇数
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = Min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

// 递归消除法，性能差距较大
// 解法参考：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/di-gui-pai-chu-fa-jian-dan-yi-dong-de-ologminmnjie/
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays1(nums2, nums1)
	}
	nums1Lens, nums2Lens := len(nums1), len(nums2)
	midM, midN := (nums1Lens-1)>>1, (nums2Lens-1)>>1

	if nums1Lens == 0 { // 处理长度为0的情况
		if nums2Lens%2 == 1 {
			return float64(nums2[midN])
		}
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}

	if nums1Lens == 1 || nums1Lens == 2 { // 边界条件
		if nums2Lens < 3 { // nums2Lens 小于 3 的情况下，取 nums2 所有元素和 nums1 的元素进行排序
			for i := 0; i < nums2Lens; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if nums2Lens&1 == 1 { // nums2Lens 大 于2 且为奇数的情况下，取 nums2 中间 3 位和 nums1 的元素进行排序
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else { // 其他情况下，取 nums2 的中间 4 位和 nums1 的元素进行排序
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
		sort.Ints(nums1)
		nums1Lens = len(nums1)
		midM = (nums1Lens - 1) >> 1

		if len(nums1)&1 == 1 {
			return float64(nums1[midM])
		} else {
			return float64(nums1[midM]+nums1[midM+1]) / 2
		}
	}

	// nums2Lens 为奇数时，midNP==midN；nums2Lens 为偶数时，midNP==midN+1。
	var midNP = midN
	if nums2Lens&1 == 0 {
		midNP++
	}

	if nums1[midM] == nums2[midNP] {
		return float64(nums1[midM])
	}
	if nums1[midM] < nums2[midNP] {
		// 消除 nums1 数组  0至 midM-1 下标的元素，和 nums2 数组 n-midM 下标之后的元素
		return findMedianSortedArrays1(nums1[midM:], nums2[:nums2Lens-midM])
	}
	// 消除 nums2 数组 0 至 midM-1 下标的元素，和 nums1 数组 n-midM 下标之后的元素
	return findMedianSortedArrays1(nums2[midM:], nums1[:nums1Lens-midM])
}
