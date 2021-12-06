/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/6 下午10:59
 * Author: beihai
 */

package Offer

/*找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
*/

func findRepeatNumber(nums []int) int {
	exist := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			return v
		}
		exist[v] = struct{}{}
	}
	return -1
}
