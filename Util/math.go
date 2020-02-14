/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 11:57 AM
 * Author: beihai
 */

package Util

func AbsInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
