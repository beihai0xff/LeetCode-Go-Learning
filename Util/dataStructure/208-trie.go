/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 24/11/21 下午2:28
 * Author: beihai
 */

package dataStructure

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, char := range word {
		char -= 'a'
		if node.children[char] == nil {
			node.children[char] = &Trie{}
		}
		node = node.children[char]
	}
	// 最后一个节点标记为 key 的结束
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefixNode(word)
	return node != nil && node.isEnd == true
}

func (t *Trie) searchPrefixNode(prefix string) *Trie {
	node := t
	for _, char := range prefix {
		char -= 'a'
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefixNode(prefix) != nil
}
