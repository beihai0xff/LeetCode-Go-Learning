/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/11 下午9:28
 * Author: beihai
 */

package _136_single_number

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：
	- 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
*/

// 异或运算的性质：任何一个数字异或它自己都等于0。从头到尾依次异或数组中的每一个数字，最终的结果就是只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

// 用一个哈希表存储每个元素出现的次数
func singleNumber1(nums []int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		numsMap[v]++
	}
	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
