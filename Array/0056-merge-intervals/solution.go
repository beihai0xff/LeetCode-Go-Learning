/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/23/20, 6:36 PM
 * Author: beihai
 */

package _056_merge_intervals

import (
	. "LeetCode/Util"
	"sort"
)

/*
给出一个区间的集合，请合并所有重叠的区间。
*/
// 将区间按起点大小排序，然后再合并
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	result = append(result, intervals[0])
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > result[curIndex][1] {
			curIndex++
			result = append(result, intervals[i])
		} else {
			result[curIndex][1] = Max(intervals[i][1], result[curIndex][1])
		}
	}
	return result
}
