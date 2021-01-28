/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/28 下午10:20
 * Author: beihai
 */

package Array

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?
*/

// 求和
func missingNumber(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res += i - nums[i]
	}
	return res + len(nums)
}
