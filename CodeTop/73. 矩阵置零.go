package CodeTop

func setZeroes(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 {
		return
	}

	setZero := func(r, c int) {
		matrix[r][c] = 0
	}

	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				for R := 0; R < row; R++ {
					defer setZero(R, j)
				}
				for C := 0; C < col; C++ {
					defer setZero(i, C)
				}
			}
		}
	}
}

func setZeroes2(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 用矩阵的第一行和第一列代替方法一中的两个标记数组，以达到 O(1)O(1) 的额外空间。
// 但这样会导致原数组的第一行和第一列被修改，无法记录它们是否原本包含 0。
// 因此我们需要额外使用两个标记变量分别记录第一行和第一列是否原本包含 0。

func setZeroes3(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
