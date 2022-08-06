package CodeTop

import "strings"

func stringMatching(words []string) []string {
	var res []string
	for i, x := range words {
		for j, y := range words {
			if j != i && strings.Contains(y, x) {
				res = append(res, x)
				break
			}
		}
	}
	return res
}
