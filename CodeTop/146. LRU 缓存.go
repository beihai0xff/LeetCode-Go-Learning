package CodeTop

import "container/list"

type LRUCache struct {
	Cap int
	// Element 中存储 pair 结构体，用哈希表指向链表的节点
	// LRUCache 执行删除操作的时候，需要维护 2 个数据结构，一个是 map，一个是双向链表。
	// 在双向链表中删除淘汰出去的 value，在 map 中删除淘汰出去 value 对应的 key。
	// 如果在双向链表的 value 中不存储 key，那么再删除 map 中的 key 的时候有点麻烦。
	Keys map[int]*list.Element
	List *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if el, ok := c.Keys[key]; ok {
		// 如果 map 中存在，将它移动到双向链表的表头
		c.List.MoveToFront(el)
		return el.Value.(pair).value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	p := pair{key: key, value: value}
	// 先查询 map 中是否存在 key
	if el, ok := c.Keys[key]; ok {
		// 如果存在，更新它的 value
		el.Value = p
		c.List.MoveToFront(el)
	} else {
		el := c.List.PushFront(p)
		c.Keys[key] = el
	}
	// 判断是否需要淘汰最后一个结点
	if c.List.Len() > c.Cap {
		el := c.List.Back()
		c.List.Remove(el)
		delete(c.Keys, el.Value.(pair).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
