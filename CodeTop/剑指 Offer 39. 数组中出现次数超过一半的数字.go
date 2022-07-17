package CodeTop

import "sort"

// 摩尔投票法
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	vote, count := nums[0], 0
	for _, num := range nums {
		if count == 0 {
			vote = num
		}
		if vote == num {
			count++
		} else {
			count--
		}
	}

	return vote
}

// 排序
func majorityElement2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sort.Ints(nums)

	return nums[len(nums)/2]
}
