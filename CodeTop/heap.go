package CodeTop

type MinHeap struct {
	arr []int
	cap int
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{
		arr: []int{},
		cap: cap,
	}
}

func (h *MinHeap) GetMin() int {
	if len(h.arr) == 0 {
		return 0
	}
	return h.arr[0]
}

func (h *MinHeap) Push(num int) {
	if len(h.arr) < h.cap {
		h.arr = append(h.arr, num)
		h.up(len(h.arr) - 1)
	} else if num > h.arr[0] {
		h.arr[0] = num
		h.down(0)
	}
}

// i 表示节点在数组中的位置
func (h *MinHeap) up(i int) {
	for {
		p := i / 2 // p 是父节点的位置
		if i == p || h.less(p, i) {
			break
		}
		h.swap(p, i)
		i = p
	}
}

// i 表示节点在数组中的位置
func (h *MinHeap) down(i int) {
	for {
		left := 2*i + 1 // 左子节点
		if left >= len(h.arr) {
			break // i已经是叶子结点了
		}
		j := left
		if right := left + 1; right < len(h.arr) && h.less(right, left) {
			j = right // 右孩子
		}
		if h.less(i, j) {
			break // 如果父节点比子节点小，则不交换
		}
		h.swap(i, j) // 交换父节点和子节点
		i = j        // 继续向下比较
	}
}

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) less(i, j int) bool {
	return h.arr[i] < h.arr[j]
}
