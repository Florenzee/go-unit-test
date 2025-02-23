package logic_1

func Soal9(n int) []int {
	result := make([]int, n)
	value := 3

	mid := n / 2 //cari index tengah

	// First half (incrementing)
	for i := 0; i < mid; i++ {
		result[i] = value
		value += 3
	}
	// Middle index (continuing +2 if n is odd)
	result[mid] = value
	if n%2 == 1 {
		value += 3
	}
	// Second half (decrementing)
	for i := mid; i < n; i++ {
		value -= 3
		result[i] = value
	}
	return result
}
