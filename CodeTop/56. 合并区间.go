package CodeTop

import "sort"

// 按照区间的左端点排序，那么在排完序的列表中，可以合并的区间一定是连续的。
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 || intervals == nil {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	prev := intervals[0]
	for fast := 1; fast < len(intervals); fast++ {
		curr := intervals[fast]
		if curr[0] > prev[1] {
			res = append(res, prev)
			prev = curr
		} else {
			prev[1] = max(curr[1], prev[1])
		}
	}

	return append(res, prev)
}
