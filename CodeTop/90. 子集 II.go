package CodeTop

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	var t []int
	var dfs func(isPreAdded bool, index int)

	dfs = func(isPreAdded bool, index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), t...))
			return
		}

		dfs(false, index+1)
		if !isPreAdded && index > 0 && nums[index-1] == nums[index] {
			return
		}
		t = append(t, nums[index])
		dfs(true, index+1)
		t = t[:len(t)-1]
	}

	dfs(false, 0)
	return res
}
