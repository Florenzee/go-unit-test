package logic_2

func Soal3(n int) (result [][]int) {
	result = make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, n)

		for j := 0; j < n; j++ {
			result[i][j] = 1 + 2*(i*9+j)
		}
	}
	return result
}
