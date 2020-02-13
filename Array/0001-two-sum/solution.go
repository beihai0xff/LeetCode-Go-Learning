package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	lens := len(nums)
	for i := 0; i < lens; i++ {
		another := target - nums[i]
		j := i
		for j < lens {
			if another == nums[j] {
				return []int{i, j}
			}
			j++
		}
	}
	return nil
}
