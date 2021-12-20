package Offer

func findContinuousSequence(target int) (res [][]int) {
	i, j, sum := 1, 2, 3

	for i <= target/2 {
		if target > sum {
			j++
			sum += j
		} else {
			if target == sum {
				tmp := make([]int, j-i+1)
				for k := i; k <= j; k++ {
					tmp[k-i] = k
				}
				res = append(res, tmp)
			}
			sum -= i
			i++
		}
	}
	return
}
