package logic_2

func Soal1(n int) (result [][]int) {
	result = make([][]int, 0)

	for i := 0; i < n; i++ {
		result[i] = make([]int, 0)
		num := 1

		for j := 0; j < i; j++ {
			result[i][j] = num
			num += 2
		}
	}
	return result
}
