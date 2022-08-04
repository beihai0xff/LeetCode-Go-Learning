package CodeTop

import "strings"

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	stack := []string{}
	for _, str := range strs {
		if str == "." || str == "" {
			continue
		}
		if str == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, str)
	}

	return "/" + strings.Join(stack, "/")
}
