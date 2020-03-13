/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/11 下午10:19
 * Author: beihai
 */

package _283_move_zeroes

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

说明:
	- 必须在原数组上操作，不能拷贝额外的数组。
	- 尽量减少操作次数。
*/

// 26、27、80、283 原理相通
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		// 当前元素 != 0，就把它交换到左边，等于 0 的交换到右边
		if nums[fast] != 0 {
			if fast != slow {
				nums[fast], nums[slow] = nums[slow], nums[fast]
			}
			slow++
		}
	}
}
