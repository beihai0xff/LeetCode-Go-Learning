package CodeTop

import "sort"

// 排序 + 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	// 枚举 a
	for first := 0; first < len(nums); first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// count 对应的指针初始指向数组的最右端
		third := len(nums) - 1
		// target = b + count = -1 * first
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < len(nums); second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 count 的指针的左侧
			// b + count > target，count 左移
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+count=0 并且 b<count 的 count 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}
