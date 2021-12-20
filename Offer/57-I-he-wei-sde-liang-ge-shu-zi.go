package Offer

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
