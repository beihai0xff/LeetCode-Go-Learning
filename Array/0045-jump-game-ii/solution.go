/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/19/20, 2:47 PM
 * Author: beihai
 */

package _045_jump_game_ii

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。
*/

// 贪心算法
func jump(nums []int) int {
	steps, position := 0, len(nums)-1
	for position != 0 {
		for i := 0; i < position; i++ {
			if nums[i] >= position-i {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
