package logic_1

func Soal4(n int) []int {
	slice := make([]int, n)
	num := 19

	for i := 0; i < n; i++ {
		slice[i] = num
		slice[i] -= 2
	}
	return slice

}
