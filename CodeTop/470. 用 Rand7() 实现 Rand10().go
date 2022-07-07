package CodeTop

func rand10() int {
	for {
		row, col := rand7(), rand7()
		pos := (row-1)*7 + col
		if pos <= 40 {
			return (pos-1)%10 + 1
		}
	}
}

func rand7() int {
	return 2
}
