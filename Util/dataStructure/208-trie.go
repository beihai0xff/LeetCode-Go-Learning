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

func (t *Trie) IsConcatenated(prefix string) bool {
	if prefix == "" {
		return false
	}
	return t.isConcatenated(prefix)
}

func (t *Trie) isConcatenated(prefix string) bool {
	if prefix == "" {
		return true
	}

	node := t
	for index, char := range prefix {
		node = node.children[char-'a']
		if node == nil {
			return false
		}
		if node.isEnd && t.isConcatenated(prefix[index+1:]) {
			return true
		}
	}
	// 走到这一步说明递归到最后到 prefix 并非是一个完整的单词
	return false
}
