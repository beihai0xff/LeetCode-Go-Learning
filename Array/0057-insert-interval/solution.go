/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/23/20, 6:59 PM
 * Author: beihai
 */

package _057_insert_interval

import . "LeetCode/Util"

/*
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	intervals = append(intervals, newInterval)
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i+1][0] > intervals[i][1] {
			// 后一个 interval 的 start 值大于前一个 interval 的 end，遍历结束
			break
		} else if intervals[i+1][1] < intervals[i][0] {
			// 后一个 interval 的 end 值小于前一个 interval 的 start 值，交换
			intervals[i], intervals[i+1] = intervals[i+1], intervals[i]
			continue
		} else {
			// 后一个 interval 与前一个有交集，则合并这两个 interval 存储于前一个 interval 中，丢弃后一个
			intervals[i][0] = Min(intervals[i][0], intervals[i+1][0])
			intervals[i][1] = Max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		}
	}
	return intervals
}
