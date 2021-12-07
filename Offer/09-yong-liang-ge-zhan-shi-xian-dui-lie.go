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
