/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/28 下午10:09
 * Author: beihai
 */

package Array

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*/

// 哈希表存储
func singleNumber(nums []int) int {
	count := map[int]int{}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for k, v := range count {
		if v == 1 {
			return k
		}
	}
	return 0
}
