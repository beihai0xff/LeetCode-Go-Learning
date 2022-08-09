package CodeTop

import (
	"container/heap"
	"sort"
)

// 哈希表排序
func topKFrequent692(words []string, k int) []string {
	if words == nil || len(words) == 0 {
		return nil
	}

	cache := map[string]int{}
	for _, v := range words {
		cache[v]++
	}

	res := make([]string, 0, len(words))

	for k, _ := range cache {
		res = append(res, k)
	}

	sort.Slice(res, func(i, j int) bool {
		if cache[res[i]] == cache[res[j]] {
			return res[i] < res[j]
		}
		return cache[res[i]] > cache[res[j]]
	})

	return res[:k]
}

// 堆排序 + 优先队列

func topKFrequent6922(words []string, k int) []string {
	if words == nil || len(words) == 0 {
		return nil
	}

	cache := map[string]int{}
	for _, v := range words {
		cache[v]++
	}

	h := &pair692Heap{}
	for key, count := range cache {
		heap.Push(h, pair692{key, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(pair692).key
	}
	return res
}

type pair692 struct {
	key   string
	count int
}

type pair692Heap []pair692

func (h pair692Heap) Len() int { return len(h) }
func (h pair692Heap) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.count < b.count || a.count == b.count && a.key > b.key
}
func (h pair692Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pair692Heap) Push(v interface{}) { *h = append(*h, v.(pair692)) }
func (h *pair692Heap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
