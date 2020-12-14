/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/14 下午2:09
 * Author: beihai
 */

package _049_group_anagrams

import "sort"

// 排序
func groupAnagrams(strs []string) [][]string {
	temp := map[string][]string{}
	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		temp[string(s)] = append(temp[string(s)], v)
	}
	res := make([][]string, 0, len(temp))
	for _, v := range temp {
		res = append(res, v)
	}
	return res
}

//
