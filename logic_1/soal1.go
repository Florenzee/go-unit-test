package logic_1

func Soal1(n int) []int {
	slice := make([]int, n)

	num := 1
	for i := 0; i < n; i++ {
		slice[i] = num
		num += 2
	}
	return slice
}
