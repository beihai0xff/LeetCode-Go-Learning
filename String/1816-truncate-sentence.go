/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/12/6 下午10:56
 * Author: beihai
 */

package String

func truncateSentence(s string, k int) string {
	i := 0
	for j, v := range s {
		if v == ' ' {
			i++
			if i == k {
				return s[:j]
			}
		}
	}
	return s
}
