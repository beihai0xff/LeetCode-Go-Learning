/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/23 下午9:10
 * Author: beihai
 */

package LeetCode_Go_Learing

/*
用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [a, b] 连接了计算机 a 和 b。

网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。

给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回 -1 。
*/

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 { // 线的数量已经不满足要求了
		return -1
	}
	graph := make([][]int, n)
	for _, c := range connections {
		v, w := c[0], c[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for _, to := range graph[from] {
			if !vis[to] {
				dfs(to)
			}
		}
	}
	var ans int
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans - 1
}
