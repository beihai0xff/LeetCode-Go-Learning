package Offer

import (
	"container/heap"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}

	h := &IntHeap{}
	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else {
			if h.Peek() > v {
				heap.Pop(h)
				heap.Push(h, v)
			}
		}
	}
	return *h
}

type IntHeap []int // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度

// Less 如果 h[i] < h[j] 生成的就是小根堆，如果 h[i] > h[j] 生成的就是大根堆
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}
