package CodeTop

import "container/heap"

// 桶排序
func topKFrequent(nums []int, n int) []int {
	var res []int
	// 计数出现次数
	hash := map[int]int{}
	for _, v := range nums {
		hash[v]++
	}

	// 出现次数为 i 的放在 list[i]
	list := make([][]int, len(nums)+1)
	for k, v := range hash {
		list[v] = append(list[v], k)
	}

	// 出现次数最多的在最后面，倒序遍历
	for i := len(list) - 1; i >= 0 && len(res) < n; i-- {
		if len(list[i]) == 0 {
			continue
		}
		res = append(res, list[i]...)
	}

	return res
}

// 堆排序
func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, freq := range occurrences {
		heap.Push(h, [2]int{key, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
