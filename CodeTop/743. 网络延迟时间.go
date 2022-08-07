package CodeTop

import "math"

const inf = math.MaxInt32

func networkDelayTime(times [][]int, n, k int) int {
	grath := make([][]int, n)
	for i := range grath {
		grath[i] = make([]int, n)
		for j := range grath[i] {
			grath[i][j] = inf
		}
	}

	for _, t := range times {
		x, y := t[0]-1, t[1]-1
		grath[x][y] = t[2]
	}

	// 当前点距离k点的最小距离
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	// 寻找当前点与k点距离的最小
	dist[k-1] = 0
	// used表示这个节点是否被使用，只有k点有最小距离才会被使用
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		x := -1
		// 寻找距离k点距离最小的节点的坐标
		for y, u := range used {
			// u为ture是表示这个节点被使用，要往后推一个
			// x == -1只是为了取到第一个点
			// distance[j] < distance[x] 目的是寻找最小点
			// 第一次就是要找到k点，之后就是在k点基础上寻找最小距离的点
			// 并且，最小路径找到后，!u可以将在k点以后未参与到最小路径的节点找出来
			// 将其加到对应的路径上，构建完整的图
			if !u && (x == -1 || dist[y] < dist[x]) {
				x = y
			}
		}
		used[x] = true
		for y, time := range grath[x] {
			dist[y] = min(dist[y], dist[x]+time)
		}
	}

	var res int
	for _, d := range dist {
		if d == inf {
			return -1
		}
		res = max(res, d)
	}
	return res
}
