package Offer

import "container/list"

type CQueue struct {
	stackHead, stackTail *list.List
}

func Constructor() CQueue {
	return CQueue{
		stackHead: list.New(),
		stackTail: list.New(),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.stackHead.PushBack(value)

}

func (c *CQueue) DeleteHead() int {
	if c.stackTail.Len() == 0 {
		for c.stackHead.Len() > 0 {
			c.stackTail.PushBack(c.stackHead.Remove(c.stackHead.Back()))
		}
	}
	if c.stackTail.Len() > 0 {
		e := c.stackTail.Back()
		c.stackTail.Remove(e)
		return e.Value.(int)
	}
	return -1
}

type CQueue1 struct {
	stackHead, stackTail []int
}

func Constructor1() CQueue1 {
	return CQueue1{
		stackHead: []int{},
		stackTail: []int{},
	}
}

func (c *CQueue1) AppendTail(value int) {
	c.stackHead = append(c.stackHead, value)
}

func (c *CQueue1) DeleteHead() int {
	if len(c.stackTail) == 0 {
		for len(c.stackHead) > 0 {
			index := len(c.stackHead) - 1
			c.stackTail = append(c.stackTail, c.stackHead[index])
			c.stackHead = c.stackHead[:index]
		}
	}
	if len(c.stackTail) > 0 {
		index := len(c.stackTail) - 1
		res := c.stackTail[index]
		c.stackTail = c.stackTail[:index]
		return res
	}
	return -1
}
