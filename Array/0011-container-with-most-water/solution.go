/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 1:19 PM
 * Author: beihai
 */

package _011_container_with_most_water

import . "LeetCode/Util"

func maxArea(height []int) int {
	max, l, r := 0, 0, len(height)-1
	for l < r {
		max = Max(max, Min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
