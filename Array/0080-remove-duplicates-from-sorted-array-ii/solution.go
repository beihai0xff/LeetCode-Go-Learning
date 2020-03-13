/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/11 下午9:48
 * Author: beihai
 */

package _080_remove_duplicates_from_sorted_array_ii

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
*/

//  26、27、80、283 原理相通，将删除的元素移动到数组后面的空间内，返回数组实际剩余的元素个数。
// 双指针法，快指针：遍历整个数组，慢指针：记录可以覆写数据的位置
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 1, 2
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
