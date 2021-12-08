package Offer

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) < 2 {
		return numbers[0]
	}
	slow, fast := 0, 1
	for fast < len(numbers) {
		if numbers[slow] > numbers[fast] {
			break
		}
		fast++
		slow++
	}
	// 旋转数组
	if fast == len(numbers) {
		return numbers[0]
	}
	return numbers[fast]
}
