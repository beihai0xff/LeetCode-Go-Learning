package CodeTop

func singleNumber137(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
	}

	for k, v := range cache {
		if v == 1 {
			return k
		}
	}

	return -1
}

// 依次确定每一个二进制位
func singleNumber137_2(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			res |= 1 << i
		}
	}
	return int(res)
}
