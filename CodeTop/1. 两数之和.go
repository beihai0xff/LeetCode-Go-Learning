package CodeTop

func twoSum(nums []int, target int) []int {
	// cache 保存着 num[i] 与 i 的映射关系
	cache := map[int]int{}
	for k, v := range nums {
		if i, ok := cache[target-v]; ok {
			return []int{i, k}
		}
		cache[v] = k
	}

	return nil
}
