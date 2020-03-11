/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/9 下午10:00
 * Author: beihai
 */

package _078_subsets

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。
*/

// 字典排序（二进制排序）
// 将每个子集映射到长度为 n 的位掩码中，其中第 i 位掩码 nums[i] 为 1，表示第 i 个元素在子集中；如果第 i 位掩码 nums[i] 为 0，表示第 i 个元素不在子集中。
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var result [][]int
	sum := 1 << uint(len(nums))
	for i := 0; i < sum; i++ {
		var stack []int
		temp := i                             // i 从 000...000 到 111...111
		for j := len(nums) - 1; j >= 0; j-- { // 遍历 i 的每一位
			if temp&1 == 1 {
				stack = append([]int{nums[j]}, stack...)
			}
			temp >>= 1
		}
		result = append(result, stack)
	}
	return result
}
