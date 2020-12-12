/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/12 下午9:48
 * Author: beihai
 */

package dataStructure

import (
	"errors"
	"sync"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 指向队列头部
	tail    int // 指向队列尾部
	lock    *sync.Mutex
}

func (c CircleQueue) Push(val int) error {
	if c.IsFull() {
		return errors.New("queue is full")
	}
	c.tail = (c.tail + 1) % c.maxSize
	c.array[c.tail] = val
	return nil
}

func (c CircleQueue) Pop() (int, error) {
	if c.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val := c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return val, nil
}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

// 队列是否为空
func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

// 队列多少个元素
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}
