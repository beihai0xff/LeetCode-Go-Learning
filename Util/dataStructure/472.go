package dataStructure

import "sort"

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	trie, res := new(Trie), []string{}

	for _, word := range words {
		if trie.IsConcatenated(word) {
			res = append(res, word)
		} else {
			trie.Insert(word)
		}
	}

	return res
}
