package Offer

import "sort"

func majorityElement(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	return nums[len(nums)/2]
}
