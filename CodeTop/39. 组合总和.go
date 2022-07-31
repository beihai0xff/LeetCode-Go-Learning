package CodeTop

func combinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	var res [][]int
	var tmp []int

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, tmp...))
			return
		} else if target < 0 {
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 或者选择当前数
		if target-candidates[idx] >= 0 {
			tmp = append(tmp, candidates[idx])
			dfs(target-candidates[idx], idx)
			tmp = tmp[:len(tmp)-1]
		}
	}

	dfs(target, 0)
	return res

}
