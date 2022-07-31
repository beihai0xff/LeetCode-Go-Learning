package CodeTop

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	cache, res := map[int]struct{}{}, 0

	for _, v := range nums {
		cache[v] = struct{}{}
	}
	for num := range cache {
		if _, ok := cache[num-1]; !ok {
			tmp := 0
			for _, ok := cache[num]; ok; _, ok = cache[num] {
				num++
				tmp++
			}
			if res < tmp {
				res = tmp
			}
		}
	}

	return res
}
