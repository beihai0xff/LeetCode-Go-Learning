package CodeTop

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	if rowIndex == 0 {
		return []int{1}
	}

	dp0 := []int{1}
	for i := 1; i <= rowIndex; i++ {
		currDP := make([]int, i+1)
		currDP[0], currDP[i] = 1, 1
		for j := 1; j < i; j++ {
			currDP[j] = dp0[j-1] + dp0[j]
		}
		dp0 = currDP
	}

	return dp0
}
