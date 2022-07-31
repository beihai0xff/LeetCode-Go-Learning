package CodeTop

import "sort"

func permuteUnique(nums []int) [][]int {
	if nums == nil && len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	visited := make([]bool, len(nums))

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] || i > 0 && !visited[i-1] && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			visited[i] = true
			dfs(path)
			visited[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs([]int{})
	return res
}
