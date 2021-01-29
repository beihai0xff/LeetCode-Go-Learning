/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/12 下午8:29
 * Author: beihai
 */

package dataStructure

import "sync"

type (
	Stack struct {
		// 栈顶节点
		top *node
		// 栈的元素数量
		len int
		// 读写锁
		lock *sync.RWMutex
	}
	node struct {
		val interface{}
		pre *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// return the len of Stack
func (s Stack) Len() int {
	return s.len
}

// View the top item on the stack
func (s Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.top.val
}

//  Push a val onto the top of the stack
func (s Stack) Push(val interface{}) {
	// 构造新节点
	n := &node{val, s.top}
	s.lock.Lock()
	s.top = n
	s.len++
	s.lock.Unlock()
}

// Pop the top element of the stack and return it
func (s Stack) Pop() interface{} {
	if s.len == 0 { // 该栈已经空了
		return nil
	}
	s.lock.Lock()
	val := s.top.val  // 取值
	s.top = s.top.pre // 弹出最顶部的节点
	s.lock.Unlock()
	return val
}

type Queue struct {
	PopStack  []int
	PushStack []int
}

/** Initialize your data structure here. */
func NewQueue() Queue {
	return Queue{
		PopStack:  []int{},
		PushStack: []int{},
	}

}

/** Push element x to the back of queue. */
func (q *Queue) Push(x int) {
	q.PushStack = append(q.PushStack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *Queue) Pop() int {
	if q.PopStack == nil { // PopStack 为空，将 PopStack 中的元素倒序移过来
		for i := len(q.PushStack) - 1; i >= 0; i-- {
			q.PopStack = append(q.PopStack, q.PushStack[i])
		}
	}
	res := q.PopStack[len(q.PopStack)-1]
	q.PopStack = q.PopStack[:len(q.PopStack)-2]
	return res
}

/** Get the front element. */
func (q *Queue) Peek() int {
	return q.PushStack[len(q.PushStack)-1]
}

/** Returns whether the queue is empty. */
func (q *Queue) Empty() bool {
	if q.PushStack == nil && q.PopStack == nil {
		return true
	}
	return false
}
