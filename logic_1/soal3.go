package logic_1

func Soal3(n int) []int {
	slice := make([]int, n)

	num := 3
	for i := 0; i < n; i++ {
		slice[i] = num
		num += 3
	}
	return slice
}
