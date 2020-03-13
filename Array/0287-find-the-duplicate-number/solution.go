/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/13 下午9:57
 * Author: beihai
 */

package _287_find_the_duplicate_number

import "sort"

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

说明：
	- 不能更改原数组（假设数组是只读的）。
	- 只能使用额外的 O(1) 的空间。
	- 时间复杂度小于 O(n2) 。
	- 数组中只有一个重复的数字，但它可能不止重复出现一次。
*/

// 解法一 快慢指针
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 相遇时不一定是重复值
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 二分搜索，遍历整个数组，统计小于等于 mid 的整数的个数，至多为 mid 个，如果超过 mid 个说明重复的数存在于区间 [low,mid] 中
func findDuplicate1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 先将数组排序，依照下标是从 0 开始递增的特性，那么数组里面的数字与下标的差值应该是越来越大。
// 如果出现了相同的数字，下标变大，差值应该比前一个数字小，出现了这个情况就说明找到了相同数字了。
func findDuplicate2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	diff := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]-i-1 >= diff {
			diff = nums[i] - i - 1
		} else {
			return nums[i]
		}
	}
	return 0
}

// 先将数组排序，判断相邻两值是否相同，原理和上面一致
func findDuplicate3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return 0
}
