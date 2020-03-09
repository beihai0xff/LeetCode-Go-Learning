/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/9 下午8:50
 * Author: beihai
 */

package _088_merge_sorted_array

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 num1 成为一个有序数组。

说明:
	- 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	- 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
*/

// 双指针，从后往前遍历，只需要循环一次
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	i := m - 1
	j := n - 1
	p := m + n - 1
	// 从后面往前放，只需要循环一次即可
	for ; i >= 0 && j >= 0; p-- {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
	}
	// 如果 nums2 中还存在元素没有放入（这些元素的值都小于 nums1[0]）
	for ; j >= 0; p-- {
		nums1[p] = nums2[j]
		j--
	}
}
