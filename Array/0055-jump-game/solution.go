/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/23/20, 6:06 PM
 * Author: beihai
 */

package _055_jump_game

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。
*/

// 利用贪心算法：如果能到达点 a，那么 a 之前也一定有一个点能够到达
func canJump(nums []int) bool {
	lastPosition := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// 如果这个位置可以跳到
		if nums[i] >= lastPosition-i {
			lastPosition = i
		}
	}
	return lastPosition == 0
}
