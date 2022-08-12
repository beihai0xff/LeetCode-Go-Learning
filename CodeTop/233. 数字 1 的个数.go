package CodeTop

// 输入一个整数n，求从1到n这n个整数的十进制表示中 x 出现的次数，x 属于 1-9。
// 例如输入12，从1到12这些整数中包含1的数字有 1、10、11、12，1一共出现了5次。
func numOfx(n, x int) int {
	if n == 0 {
		return 0
	}

	var res int
	for i := 1; i <= n; i++ {
		for j := i; j > 0; j /= 10 {
			if j%10 == x {
				res++
			}
		}
	}

	return res
}

// 计算每一位1出现的次数
func numOfx2(n, x int) int {
	// i表示n的位数,每次进位
	// 从个位开始
	// 1234567为例,计算每一位出现1的次数
	count := 0
	for i := 1; i <= n; i *= 10 {
		// 十位的1被循环的次数
		// a = 123456
		a := n / (i * 10)
		// b = 7
		b := n % (i * 10)
		// 个位大于 1234560,出现1的次数
		// 1
		c := min(max(b-i+x, 0), i)
		// 123456 + 1
		count += a*i + c
	}
	return count
}

// 计算每一位1出现的次数
func countDigitOne(n int) int {
	// i表示n的位数,每次进位
	// 从个位开始
	// 1234567为例,计算每一位出现1的次数
	count := 0
	for i := 1; i <= n; i *= 10 {
		// 十位的1被循环的次数
		// a = 123456
		a := n / (i * 10)
		// b = 7
		b := n % (i * 10)
		// 个位大于 1234560,出现1的次数
		// 1
		c := min(max(b-i+1, 0), i)
		// 123456 + 1
		count += a*i + c
	}
	return count
}
